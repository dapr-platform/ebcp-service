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
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("player init panic: %v", err)
		}
	}()
	go maintainConnections()
	go updatePlayerPrograms()
	common.RegisterUpsertBeforeHook(model.Ebcp_playerTableInfo.Name, UpsertPlayer)
}

func UpsertPlayer(r *http.Request, in any) (out any, err error) {
	playerMu.Lock()
	defer playerMu.Unlock()
	player := in.(model.Ebcp_player)
	id := player.ID
	address := player.IPAddress + ":" + cast.ToString(player.Port)

	// 如果已存在客户端连接，先关闭它，防止资源泄漏
	if oldClient, exists := playerClients[id]; exists {
		oldClient.Close()
	}

	client, err := client.NewTCPClient(address)
	if err != nil {
		common.Logger.Errorf("Failed to connect to player %s at %s: %v", id, address, err)
		return nil, err
	}
	playerClients[id] = client
	return player, nil
}

func GetPlayerClient(id string) *client.PlayerClient {
	playerMu.RLock()
	defer playerMu.RUnlock()
	return playerClients[id]
}

// maintainConnections periodically checks and maintains player connections
func maintainConnections() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		checkConnections()
	}
}

// updatePlayerPrograms periodically updates program list for each player
func updatePlayerPrograms() {
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
				programs, err := cli.GetProgramList()
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 节目列表失败: %v", id, err)
					if client.IsTimeoutError(err) {
						player.Status = 2
						common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
					}
					return
				}

				common.Logger.Infof("播放器 [%s] 获取到 %d 个节目", id, len(programs.Programs))

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
				firstProgramId := ""

				// Insert new programs
				for _, program := range programs.Programs {
					if program.IsEmpty {
						continue
					}
					if firstProgramId == "" {
						firstProgramId = cast.ToString(program.ID)
					}

					programIdStr := cast.ToString(program.ID)
					playerProgram := model.Ebcp_player_program{
						ID:          common.GetMD5Hash(id + "_" + programIdStr),
						CreatedBy:   "admin",
						CreatedTime: common.LocalTime(time.Now()),
						UpdatedBy:   "admin",
						UpdatedTime: common.LocalTime(time.Now()),
						ProgramID:   programIdStr,
						Name:        program.Name,
						PlayerID:    id,
					}

					// 检查是否为新增或更新
					isNewProgram := true

					for _, existingProgram := range exists {
						if existingProgram.ProgramID == programIdStr {
							isNewProgram = false
							break
						}
					}

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

						syncPlayerProgramMedia(cli, id, cast.ToString(playerProgram.ID), program.ID)
						addedMap[programIdStr] = true
					}
				}

				// Delete programs that are no longer in the program list
				deleteCount := 0
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

				if firstProgramId != "" {
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
					if player.CurrentProgramID == "" {
						common.Logger.Infof("初始化播放器 [%s] 当前播放节目为 [%s]", id, firstProgramId)
						player.CurrentProgramID = firstProgramId
						err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
						if err != nil {
							common.Logger.Errorf("更新播放器 [%s] 当前播放节目失败: %v", id, err)
						}
					}
				}

				currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 当前播放节目失败: %v", id, err)

					return
				}
				player.CurrentProgramID = currentProgramId
				player.CurrentProgramState = cast.ToInt32(currentProgramState)
				err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
				if err != nil {
					common.Logger.Errorf("更新播放器 [%s] 当前播放节目失败: %v", id, err)
				}
			}(playerId, playerClient)
		}

		common.Logger.Info("播放器程序更新任务已启动，等待所有更新完成")
	}
}
func GetPlayerCurrentProgram(client *client.PlayerClient) (programId string, state int, err error) {
	currentProgram, err := client.GetCurrentProgram()
	if err != nil {
		common.Logger.Errorf("Failed to get current program: %v", err)
		return
	}
	programId = cast.ToString(currentProgram.ProgramID)
	state = int(currentProgram.ProgState)
	return
}
func syncPlayerProgramMedia(client *client.PlayerClient, playerId, programId string, playerProgramID uint32) {
	common.Logger.Infof("Syncing program media for player %s program %s", playerId, programId)
	medias, err := client.GetAllProgramMedia(playerProgramID)
	if err != nil {
		common.Logger.Errorf("Failed to get program media list from player %s, program ID: %d: %v",
			playerId, playerProgramID, err)
		return
	}
	exists, err := common.DbQuery[model.Ebcp_player_program_media](context.Background(), common.GetDaprClient(),
		model.Ebcp_player_program_mediaTableInfo.Name, "program_id="+programId)
	if err != nil {
		common.Logger.Errorf("Failed to check if player %s has program media: %v", playerId, err)
		return
	}
	addedMap := make(map[string]bool)
	for _, media := range medias.Media {
		playerProgramMedia := model.Ebcp_player_program_media{
			ID:              common.GetMD5Hash(playerId + "_" + programId + "_" + cast.ToString(media.ID)),
			CreatedBy:       "admin",
			CreatedTime:     common.LocalTime(time.Now()),
			UpdatedBy:       "admin",
			UpdatedTime:     common.LocalTime(time.Now()),
			MediaID:         cast.ToString(media.ID),
			MediaName:       media.Name,
			PlayerID:        playerId,
			ProgramID:       programId,
			PlayerProgramID: cast.ToString(playerProgramID),
		}
		err = common.DbUpsert[model.Ebcp_player_program_media](context.Background(), common.GetDaprClient(),
			playerProgramMedia, model.Ebcp_player_program_mediaTableInfo.Name, "id")
		if err != nil {
			common.Logger.Errorf("Failed to upsert program media for player %s program %s: %v", playerId, programId, err)
		} else {
			addedMap[cast.ToString(media.ID)] = true
		}
	}
	for _, media := range exists {
		if !addedMap[cast.ToString(media.MediaID)] {
			err = common.DbDelete(context.Background(), common.GetDaprClient(),
				model.Ebcp_player_program_mediaTableInfo.Name, "id", media.ID)
			if err != nil {
				common.Logger.Errorf("Failed to delete program media for player %s program %s: %v", playerId, programId, err)
			}
		}
	}
	common.Logger.Infof("Synced program media for player %s program %s", playerId, programId)
}

func GetPlayerPrograms(playerId string) ([]model.Ebcp_player_program_info, error) {
	programs, err := common.DbQuery[model.Ebcp_player_program_info](context.Background(), common.GetDaprClient(), model.Ebcp_player_program_infoTableInfo.Name, "player_id="+playerId)
	if err != nil {
		return nil, err
	}
	return programs, nil
}
func PlayerPlay(player *model.Ebcp_player) error {
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStatePlay {
		err = playerClient.PlayProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("播放播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStatePlay
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	}
	return nil
}
func PlayerPause(player *model.Ebcp_player) error {
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStatePause {
		err = playerClient.PauseProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("暂停播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStatePause
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	}
	return nil
}
func PlayerStop(player *model.Ebcp_player) error {
	playerClient := GetPlayerClient(player.ID)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	currentProgramId, currentProgramState, err := GetPlayerCurrentProgram(playerClient)
	if err != nil {
		return fmt.Errorf("获取播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}
	player.CurrentProgramID = currentProgramId
	player.CurrentProgramState = cast.ToInt32(currentProgramState)
	if player.CurrentProgramState != ProgramStateStop {
		err = playerClient.StopProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			return fmt.Errorf("停止播放器 [%s] 节目失败: %v", player.ID, err)
		}
		player.CurrentProgramState = ProgramStateStop
		err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
		if err != nil {
			return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
		}
	}
	return nil
}
func PlayProgram(playerId, programId string) error {
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
	err = playerClient.PlayProgram(cast.ToUint32(programId))
	if err != nil {
		return fmt.Errorf("播放播放器 [%s] 节目失败: %v", playerId, err)
	}

	player.CurrentProgramID = programId
	player.CurrentProgramState = ProgramStatePlay
	err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新播放器 [%s] 当前播放节目失败: %v", player.ID, err)
	}

	return nil
}
func PauseProgram(playerId, programId string) error {
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
func SetProgramMediaProcess(playerId, programId, mediaId string, remainTime, totalTime int) error {
	playerClient := GetPlayerClient(playerId)
	if playerClient == nil {
		return fmt.Errorf("player not found")
	}
	err := playerClient.ControlLayerProgress(programId, uint32(remainTime), uint32(totalTime), cast.ToUint16(mediaId))
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
