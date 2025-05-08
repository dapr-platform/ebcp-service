package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"ebcp-service/client"
	"ebcp-service/model"

	"github.com/dapr-platform/common"
	"github.com/spf13/cast"
)

const (
	PlayerStatusOnline  = 1
	PlayerStatusOffline = 2
	PlayerStatusError   = 3
	ProgramStatePlay    = 0
	ProgramStatePause   = 1
	ProgramStateStop    = 2
)

var (
	playerClients = make(map[string]*client.PlayerClient)
	playerMu      sync.RWMutex
	// 新增播放器状态锁映射，用于同步对单个播放器状态的修改
	playerStateLocks = make(map[string]*sync.Mutex)
	playerStateMu    sync.RWMutex
)

// 获取播放器状态锁
func getPlayerStateLock(playerId string) *sync.Mutex {
	playerStateMu.RLock()
	lock, exists := playerStateLocks[playerId]
	playerStateMu.RUnlock()

	if !exists {
		playerStateMu.Lock()
		lock, exists = playerStateLocks[playerId]
		if !exists {
			lock = &sync.Mutex{}
			playerStateLocks[playerId] = lock
		}
		playerStateMu.Unlock()
	}

	return lock
}

func init() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("player init panic: %v", err)
		}
	}()
	common.Logger.Info("初始化播放器服务...")
	go maintainConnections()
	go updatePlayerPrograms()
	common.RegisterUpsertBeforeHook(model.Ebcp_playerTableInfo.Name, UpsertPlayer)
	common.Logger.Info("播放器服务初始化完成")
}

func UpsertPlayer(r *http.Request, in any) (out any, err error) {
	common.Logger.Infof("开始更新/插入播放器...")
	playerMu.Lock()
	defer playerMu.Unlock()
	player := in.(model.Ebcp_player)
	id := player.ID
	address := player.IPAddress + ":" + cast.ToString(player.Port)
	common.Logger.Infof("处理播放器 [%s] 地址: %s", id, address)

	// 如果已存在客户端连接，先关闭它，防止资源泄漏
	if oldClient, exists := playerClients[id]; exists {
		common.Logger.Infof("关闭播放器 [%s] 的现有连接", id)
		oldClient.Close()
	}

	common.Logger.Infof("尝试连接播放器 [%s] 地址: %s", id, address)
	client, err := client.NewTCPClient(address)
	if err != nil {
		common.Logger.Errorf("Failed to connect to player %s at %s: %v", id, address, err)
		return nil, err
	}
	common.Logger.Infof("成功连接到播放器 [%s]", id)
	playerClients[id] = client
	return player, nil
}

func GetPlayerClient(id string) *client.PlayerClient {
	common.Logger.Debugf("获取播放器 [%s] 客户端", id)
	playerMu.RLock()
	defer playerMu.RUnlock()
	client := playerClients[id]
	if client == nil {
		common.Logger.Warnf("播放器 [%s] 客户端不存在", id)
	}
	return client
}

// maintainConnections periodically checks and maintains player connections
func maintainConnections() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("maintainConnections panic: %v", err)
			// 重启该goroutine
			go maintainConnections()
		}
	}()
	common.Logger.Info("启动播放器连接维护定时任务")
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		common.Logger.Debug("开始检查播放器连接状态")
		checkConnections()
		common.Logger.Debug("播放器连接状态检查完成")
	}
}

// updatePlayerPrograms periodically updates program list for each player
func updatePlayerPrograms() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("updatePlayerPrograms panic: %v", err)
			// 重启该goroutine
			go updatePlayerPrograms()
		}
	}()
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	common.Logger.Info("启动播放器程序更新定时任务")

	for range ticker.C {
		common.Logger.Info("开始执行播放器程序更新")

		// 复制当前客户端映射，避免长时间持有锁
		clientsCopy := make(map[string]*client.PlayerClient)
		playerMu.RLock()
		for id, client := range playerClients {
			clientsCopy[id] = client
		}
		playerMu.RUnlock()

		playerCount := len(clientsCopy)
		common.Logger.Infof("当前共有 %d 个播放器需要更新", playerCount)

		if playerCount == 0 {
			common.Logger.Info("没有可用的播放器，跳过本次更新")
			continue
		}

		// 使用复制的映射处理程序更新
		for playerId, playerClient := range clientsCopy {
			go func(id string, cli *client.PlayerClient) {
				defer func() {
					if err := recover(); err != nil {
						common.Logger.Errorf("updatePlayerPrograms goroutine for player %s panic: %v", id, err)
					}
				}()
				common.Logger.Infof("开始更新播放器 [%s] 的节目列表", id)
				player, err := common.DbGetOne[model.Ebcp_player](context.Background(), common.GetDaprClient(),
					model.Ebcp_playerTableInfo.Name, "id="+id)
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 信息失败: %v", id, err)
					return
				}
				if player == nil {
					common.Logger.Errorf("播放器 [%s] 不存在", id)
					return
				}
				common.Logger.Debugf("开始从播放器 [%s] 获取节目列表", id)
				programs, err := cli.GetProgramList()
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 节目列表失败: %v", id, err)
					if client.IsTimeoutError(err) {
						common.Logger.Warnf("播放器 [%s] 连接超时，标记为错误状态", id)
						player.Status = PlayerStatusError
						common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
					}
					return
				}

				// 获取播放器状态锁
				stateLock := getPlayerStateLock(id)
				stateLock.Lock()

				common.Logger.Debugf("获取播放器 [%s] 当前播放节目信息", id)
				currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 当前播放节目失败: %v", id, err)
					stateLock.Unlock()
					return
				}
				common.Logger.Infof("播放器 [%s] 当前播放节目ID: %s, 状态: %d", id, currentProgramId, currentProgramState)
				player.CurrentProgramID = currentProgramId
				player.CurrentProgramState = cast.ToInt32(currentProgramState)

				volume, muteState, err := GetPlayerVolumeAndMuteState(playerClient)
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 音量和静音状态失败: %v", id, err)
				} else {
					player.Volume = int32(volume)
					player.SoundState = int32(muteState)
				}
				err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
				if err != nil {
					common.Logger.Errorf("更新播放器 [%s] 当前播放节目失败: %v", id, err)
				} else {
					common.Logger.Debugf("成功更新播放器 [%s] 当前播放节目信息", id)
				}

				stateLock.Unlock()

				common.Logger.Infof("播放器 [%s] 获取到 %d 个节目", id, len(programs.Programs))

				common.Logger.Debugf("查询播放器 [%s] 数据库中已有节目", id)
				exists, err := common.DbQuery[model.Ebcp_player_program](context.Background(), common.GetDaprClient(),
					model.Ebcp_player_programTableInfo.Name, "player_id="+id)
				if err != nil {
					common.Logger.Errorf("查询播放器 [%s] 已有节目失败: %v", id, err)
					return
				}

				common.Logger.Infof("播放器 [%s] 数据库中有 %d 个节目记录", id, len(exists))

				addedMap := make(map[string]bool)
				updatedCount := 0
				addedCount := 0

				// Insert new programs
				common.Logger.Debugf("开始同步播放器 [%s] 节目列表", id)
				for _, program := range programs.Programs {
					if program.IsEmpty {
						common.Logger.Debugf("跳过空节目记录")
						continue
					}

					var state = int32(ProgramStateStop)
					programIdStr := cast.ToString(program.ID)
					if programIdStr == player.CurrentProgramID {
						state = player.CurrentProgramState
						common.Logger.Debugf("节目 [%s] 是当前播放节目，状态为: %d", programIdStr, state)
					}
					playerProgram := model.Ebcp_player_program{
						ID:          common.GetMD5Hash(id + "_" + programIdStr),
						CreatedBy:   "admin",
						CreatedTime: common.LocalTime(time.Now()),
						UpdatedBy:   "admin",
						UpdatedTime: common.LocalTime(time.Now()),
						ProgramID:   programIdStr,
						Name:        program.Name,
						PlayerID:    id,
						State:       state,
					}

					// 检查是否为新增或更新
					isNewProgram := true

					for _, existingProgram := range exists {
						if existingProgram.ProgramID == programIdStr {
							isNewProgram = false
							break
						}
					}

					common.Logger.Debugf("准备更新/插入节目 [%s:%s]", programIdStr, program.Name)
					err = common.DbUpsert[model.Ebcp_player_program](context.Background(), common.GetDaprClient(),
						playerProgram, model.Ebcp_player_programTableInfo.Name, "id")
					if err != nil {
						common.Logger.Errorf("更新/插入播放器 [%s] 节目 [%s:%s] 失败: %v", id, programIdStr, program.Name, err)
					} else {
						if isNewProgram {
							addedCount++
							common.Logger.Infof("为播放器 [%s] 新增节目 [%s:%s]", id, programIdStr, program.Name)
						} else {
							updatedCount++
							common.Logger.Debugf("更新播放器 [%s] 节目 [%s:%s]", id, programIdStr, program.Name)
						}

						common.Logger.Debugf("开始同步播放器 [%s] 节目 [%s] 的媒体文件", id, programIdStr)
						syncPlayerProgramMedia(cli, id, cast.ToString(playerProgram.ID), program.ID)
						addedMap[programIdStr] = true
					}
				}

				// Delete programs that are no longer in the program list
				deleteCount := 0
				common.Logger.Debugf("检查需要删除的节目")
				for _, program := range exists {
					if !addedMap[cast.ToString(program.ProgramID)] {
						common.Logger.Infof("删除播放器 [%s] 中不存在的节目 [%s:%s]", id, program.ProgramID, program.Name)
						err = common.DbDelete(context.Background(), common.GetDaprClient(),
							model.Ebcp_player_programTableInfo.Name, "id", program.ID)
						if err != nil {
							common.Logger.Errorf("删除播放器 [%s] 节目 [%s:%s] 失败: %v", id, program.ProgramID, program.Name, err)
						} else {
							deleteCount++
						}
					}
				}

				common.Logger.Infof("播放器 [%s] 节目同步完成: 新增 %d 个, 更新 %d 个, 删除 %d 个",
					id, addedCount, updatedCount, deleteCount)

			}(playerId, playerClient)
		}

		common.Logger.Info("播放器程序更新任务已启动，等待所有更新完成")
	}
}
func GetPlayerVolumeAndMuteState(client *client.PlayerClient) (volume int, muteState int, err error) {
	common.Logger.Debug("获取播放器音量和静音状态")
	resp, err := client.QueryLayerVolumeMuteStatus(0)
	if err != nil {
		common.Logger.Errorf("获取播放器音量和静音状态失败: %v", err)
		return
	}
	common.Logger.Debugf("播放器音量: %d, 静音状态: %d", resp.Volume, resp.MuteFlag)
	return int(resp.Volume), int(resp.MuteFlag), nil
}
func GetPlayerCurrentProgram(client *client.PlayerClient) (programId string, state int, err error) {
	common.Logger.Debug("获取当前播放节目信息")
	currentProgram, err := client.GetCurrentProgram()
	if err != nil {
		common.Logger.Errorf("Failed to get current program: %v", err)
		return
	}
	programId = cast.ToString(currentProgram.ProgramID)
	state = int(currentProgram.ProgState)
	common.Logger.Debugf("当前播放节目ID: %s, 状态: %d", programId, state)
	return
}
func syncPlayerProgramMedia(client *client.PlayerClient, playerId, programId string, playerProgramID uint32) {
	common.Logger.Infof("Syncing program media for player %s program %s", playerId, programId)
	common.Logger.Debugf("从播放器 [%s] 获取节目 [%d] 的媒体列表", playerId, playerProgramID)
	medias, err := client.GetAllProgramMedia(playerProgramID)
	if err != nil {
		common.Logger.Errorf("Failed to get program media list from player %s, program ID: %d: %v",
			playerId, playerProgramID, err)
		return
	}
	common.Logger.Debugf("播放器 [%s] 节目 [%d] 有 %d 个媒体文件", playerId, playerProgramID, len(medias.Media))

	common.Logger.Debugf("查询数据库中播放器 [%s] 节目 [%s] 已有媒体文件", playerId, programId)
	exists, err := common.DbQuery[model.Ebcp_player_program_media](context.Background(), common.GetDaprClient(),
		model.Ebcp_player_program_mediaTableInfo.Name, "program_id="+programId)
	if err != nil {
		common.Logger.Errorf("Failed to check if player %s has program media: %v", playerId, err)
		return
	}
	common.Logger.Debugf("数据库中播放器 [%s] 节目 [%s] 有 %d 个媒体文件记录", playerId, programId, len(exists))

	addedMap := make(map[string]bool)
	addedCount := 0
	updatedCount := 0

	for _, media := range medias.Media {
		mediaIdStr := cast.ToString(media.ID)
		common.Logger.Debugf("处理媒体文件 [%s:%s]", mediaIdStr, media.Name)

		playerProgramMedia := model.Ebcp_player_program_media{
			ID:              common.GetMD5Hash(playerId + "_" + programId + "_" + mediaIdStr),
			CreatedBy:       "admin",
			CreatedTime:     common.LocalTime(time.Now()),
			UpdatedBy:       "admin",
			UpdatedTime:     common.LocalTime(time.Now()),
			MediaID:         mediaIdStr,
			MediaName:       media.Name,
			PlayerID:        playerId,
			ProgramID:       programId,
			PlayerProgramID: cast.ToString(playerProgramID),
		}

		// 检查是否为新增或更新
		isNewMedia := true
		for _, existingMedia := range exists {
			if existingMedia.MediaID == mediaIdStr {
				isNewMedia = false
				break
			}
		}

		err = common.DbUpsert[model.Ebcp_player_program_media](context.Background(), common.GetDaprClient(),
			playerProgramMedia, model.Ebcp_player_program_mediaTableInfo.Name, "id")
		if err != nil {
			common.Logger.Errorf("Failed to upsert program media for player %s program %s: %v", playerId, programId, err)
		} else {
			if isNewMedia {
				addedCount++
				common.Logger.Debugf("为播放器 [%s] 节目 [%s] 新增媒体 [%s:%s]", playerId, programId, mediaIdStr, media.Name)
			} else {
				updatedCount++
				common.Logger.Debugf("更新播放器 [%s] 节目 [%s] 媒体 [%s:%s]", playerId, programId, mediaIdStr, media.Name)
			}
			addedMap[mediaIdStr] = true
		}
	}

	deleteCount := 0
	for _, media := range exists {
		if !addedMap[media.MediaID] {
			common.Logger.Debugf("删除播放器 [%s] 节目 [%s] 中不存在的媒体 [%s:%s]", playerId, programId, media.MediaID, media.MediaName)
			err = common.DbDelete(context.Background(), common.GetDaprClient(),
				model.Ebcp_player_program_mediaTableInfo.Name, "id", media.ID)
			if err != nil {
				common.Logger.Errorf("Failed to delete program media for player %s program %s: %v", playerId, programId, err)
			} else {
				deleteCount++
			}
		}
	}

	common.Logger.Infof("播放器 [%s] 节目 [%s] 媒体同步完成: 新增 %d 个, 更新 %d 个, 删除 %d 个",
		playerId, programId, addedCount, updatedCount, deleteCount)
}

func GetPlayerPrograms(playerId string) ([]model.Ebcp_player_program_info, error) {
	common.Logger.Debugf("获取播放器 [%s] 节目列表", playerId)
	programs, err := common.DbQuery[model.Ebcp_player_program_info](context.Background(), common.GetDaprClient(), model.Ebcp_player_program_infoTableInfo.Name, "player_id="+playerId)
	if err != nil {
		common.Logger.Errorf("获取播放器 [%s] 节目列表失败: %v", playerId, err)
		return nil, err
	}
	common.Logger.Debugf("播放器 [%s] 有 %d 个节目", playerId, len(programs))
	return programs, nil
}
func PlayerPlay(player *model.Ebcp_player) error {
	common.Logger.Infof("播放器 [%s] 开始播放", player.ID)
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(player.ID)
	stateLock.Lock()
	defer stateLock.Unlock()

	common.Logger.Debugf("获取播放器 [%s] 当前播放节目", player.ID)
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	common.Logger.Infof("播放器 [%s] 当前节目ID: %s, 状态: %d", player.ID, currentProgramId, currentProgramState)
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStatePlay {
		common.Logger.Infof("播放器 [%s] 当前节目状态不是播放状态，发送播放命令", player.ID)
		err = playerClient.PlayProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("播放播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStatePlay
		common.Logger.Debugf("更新播放器 [%s] 状态为播放状态", player.ID)
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	} else {
		common.Logger.Infof("播放器 [%s] 当前节目已经是播放状态", player.ID)
	}
	return updatePlayerProgramState(player, ProgramStatePlay)
}

func updatePlayerProgramState(player *model.Ebcp_player, state int32) error {
	common.Logger.Debugf("更新播放器 [%s] 所有节目状态", player.ID)
	programs, err := common.DbQuery[model.Ebcp_player_program](context.Background(), common.GetDaprClient(), model.Ebcp_player_programTableInfo.Name, "player_id="+player.ID)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 节目列表失败: %v", player.ID, err)
	}
	common.Logger.Debugf("播放器 [%s] 有 %d 个节目需要更新状态", player.ID, len(programs))
	for _, program := range programs {
		if program.ProgramID == player.CurrentProgramID {
			common.Logger.Debugf("更新当前播放节目 [%s] 状态为 %d", program.ProgramID, state)
			program.State = state
		} else {
			common.Logger.Debugf("更新非当前播放节目 [%s] 状态为停止", program.ProgramID)
			program.State = ProgramStateStop
		}
		err = common.DbUpsert[model.Ebcp_player_program](context.Background(), common.GetDaprClient(), program, model.Ebcp_player_programTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 节目状态失败: %v", player.ID, err)
		}
	}
	common.Logger.Infof("播放器 [%s] 所有节目状态更新完成", player.ID)
	return nil
}

func PlayerPause(player *model.Ebcp_player) error {
	common.Logger.Infof("播放器 [%s] 暂停播放", player.ID)
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(player.ID)
	stateLock.Lock()
	defer stateLock.Unlock()

	common.Logger.Debugf("获取播放器 [%s] 当前播放节目", player.ID)
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	common.Logger.Infof("播放器 [%s] 当前节目ID: %s, 状态: %d", player.ID, currentProgramId, currentProgramState)
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStatePause {
		common.Logger.Infof("播放器 [%s] 当前节目状态不是暂停状态，发送暂停命令", player.ID)
		err = playerClient.PauseProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("暂停播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStatePause
		common.Logger.Debugf("更新播放器 [%s] 状态为暂停状态", player.ID)
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	} else {
		common.Logger.Infof("播放器 [%s] 当前节目已经是暂停状态", player.ID)
	}
	return updatePlayerProgramState(player, ProgramStatePause)
}
func PlayerStop(player *model.Ebcp_player) error {
	common.Logger.Infof("播放器 [%s] 停止播放", player.ID)
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(player.ID)
	stateLock.Lock()
	defer stateLock.Unlock()

	common.Logger.Debugf("获取播放器 [%s] 当前播放节目", player.ID)
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	common.Logger.Infof("播放器 [%s] 当前节目ID: %s, 状态: %d", player.ID, currentProgramId, currentProgramState)
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStateStop {
		common.Logger.Infof("播放器 [%s] 当前节目状态不是停止状态，发送停止命令", player.ID)
		err = playerClient.StopProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("停止播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStateStop
		common.Logger.Debugf("更新播放器 [%s] 状态为停止状态", player.ID)
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	} else {
		common.Logger.Infof("播放器 [%s] 当前节目已经是停止状态", player.ID)
	}
	return updatePlayerProgramState(player, ProgramStateStop)
}
func PlayProgram(playerId, programId string) error {
	common.Logger.Infof("播放器 [%s] 播放节目 [%s]", playerId, programId)
	player, err := common.DbGetOne[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "id="+playerId)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 信息失败: %v", playerId, err)
	}
	if player == nil {
		return fmt.Errorf("播放器 [%s] 不存在", playerId)
	}
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(playerId)
	stateLock.Lock()
	defer stateLock.Unlock()

	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", playerId, err)
	}
	common.Logger.Infof("播放器 [%s] 当前节目ID: %s, 状态: %d", playerId, currentProgramId, currentProgramState)
	if currentProgramId == programId && currentProgramState == ProgramStatePlay {
		common.Logger.Infof("播放器 [%s] 当前节目已经是 [%s], 无需重复播放", playerId, programId)
		return nil
	}
	if currentProgramId != programId {
		if currentProgramState == ProgramStatePlay {
			common.Logger.Infof("播放器 [%s] 当前节目不是 [%s], 发送停止命令", playerId, programId)
			err = playerClient.StopProgram(cast.ToUint32(currentProgramId))
			if err != nil {
				return fmt.Errorf("停止播放器 [%s] 节目失败: %v", playerId, err)
			}
		}
	}
	common.Logger.Debugf("向播放器 [%s] 发送播放节目 [%s] 命令", playerId, programId)
	err = playerClient.PlayProgram(cast.ToUint32(programId))
	if err != nil {
		return fmt.Errorf("播放播放器 [%s] 节目失败: %v", playerId, err)
	}

	common.Logger.Debugf("更新播放器 [%s] 当前播放节目为 [%s]", playerId, programId)
	player.CurrentProgramID = programId
	player.CurrentProgramState = ProgramStatePlay
	err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}

	return nil
}
func PauseProgram(playerId, programId string) error {
	common.Logger.Infof("播放器 [%s] 暂停节目 [%s]", playerId, programId)
	player, err := common.DbGetOne[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "id="+playerId)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 信息失败: %v", playerId, err)
	}
	if player == nil {
		return fmt.Errorf("播放器 [%s] 不存在", playerId)
	}
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(playerId)
	stateLock.Lock()
	defer stateLock.Unlock()

	err = playerClient.PauseProgram(cast.ToUint32(programId))
	if err != nil {
		return fmt.Errorf("暂停播放器 [%s] 节目失败: %v", playerId, err)
	}
	player.CurrentProgramState = ProgramStatePause
	err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	return nil
}
func StopProgram(playerId, programId string) error {
	player, err := common.DbGetOne[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "id="+playerId)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 信息失败: %v", playerId, err)
	}
	if player == nil {
		return fmt.Errorf("播放器 [%s] 不存在", playerId)
	}
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}

	// 获取播放器状态锁
	stateLock := getPlayerStateLock(playerId)
	stateLock.Lock()
	defer stateLock.Unlock()

	err = playerClient.StopProgram(cast.ToUint32(programId))
	if err != nil {
		return fmt.Errorf("停止播放器 [%s] 节目失败: %v", playerId, err)
	}
	player.CurrentProgramState = ProgramStateStop
	err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	return nil
}
func PlayProgramMedia(playerId, programId, mediaId string) error {
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	err := playerClient.PlayLayerMedia(cast.ToUint32(mediaId))
	if err != nil {
		return fmt.Errorf("播放播放器 [%s] 节目媒体失败: %v", playerId, err)
	}
	return nil
}

func PauseProgramMedia(playerId, programId, mediaId string) error {
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	err := playerClient.PauseLayerMedia(cast.ToUint32(mediaId))
	if err != nil {
		return fmt.Errorf("暂停播放器 [%s] 节目媒体失败: %v", playerId, err)
	}
	return nil
}

func GetProgramMediaProcess(playerId, programId, mediaId string) (int, int, error) {
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return 0, 0, fmt.Errorf("player not found")
	}
	progress, err := playerClient.QueryLayerProgress(cast.ToUint16(mediaId))
	if err != nil {
		return 0, 0, err
	}
	return int(progress.RemainTime), int(progress.TotalTime), nil
}
func SetProgramMediaProcess(playerId, programId, mediaId string, currentTime, totalTime int) error {
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	err := playerClient.ControlLayerProgress(programId, uint32(currentTime), uint32(totalTime), cast.ToUint16(mediaId))
	if err != nil {
		return fmt.Errorf("设置播放器 [%s] 节目媒体进度失败: %v", playerId, err)
	}
	return nil
}

// checkConnections verifies and refreshes player connections
func checkConnections() {
	// 先查询数据库，避免长时间持有锁
	players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "")
	if err != nil {
		common.Logger.Errorf("Failed to query players from database: %v", err)
		return
	}

	playerMu.Lock()
	defer playerMu.Unlock()

	// Update player clients
	for _, player := range players {
		id := player.ID
		address := player.IPAddress + ":" + cast.ToString(player.Port)

		if oldClient, exists := playerClients[id]; exists {
			// 检查地址是否变更，如果变更则更新连接
			oldAddress := oldClient.GetAddress()
			if oldAddress != address {
				oldClient.Close() // 关闭旧连接

				// 创建新连接
				client, err := client.NewTCPClient(address)
				if err != nil {
					common.Logger.Errorf("Failed to connect to player %s at %s: %v", id, address, err)
					continue
				}
				playerClients[id] = client
			}
		} else {
			// 创建新连接
			client, err := client.NewTCPClient(address)
			if err != nil {
				common.Logger.Errorf("Failed to connect to player %s at %s: %v", id, address, err)
				continue
			}
			playerClients[id] = client
		}
	}

	// 移除不存在于数据库的播放器连接
	for id, client := range playerClients {
		found := false
		for _, player := range players {
			if player.ID == id {
				found = true
				break
			}
		}
		if !found {
			client.Close()
			delete(playerClients, id)
		}
	}
}
