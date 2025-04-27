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

var (
	playerClients = make(map[string]*client.PlayerClient)
	playerMu      sync.RWMutex
)

func init() {
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
			go func(id string, client *client.PlayerClient) {
				common.Logger.Infof("开始更新播放器 [%s] 的节目列表", id)

				programs, err := client.GetProgramList()
				if err != nil {
					common.Logger.Errorf("获取播放器 [%s] 节目列表失败: %v", id, err)
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

						syncPlayerProgramMedia(client, id, cast.ToString(playerProgram.ID), program.ID)
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

			}(playerId, playerClient)
		}

		common.Logger.Info("播放器程序更新任务已启动，等待所有更新完成")
	}
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
func PlayProgramMedia(playerId, programId, mediaId string) error {
	player := GetPlayerClient(playerId)
	if player == nil {
		return fmt.Errorf("player not found")
	}
	return player.PlayLayerMedia(cast.ToUint32(mediaId))
}

func PauseProgramMedia(playerId, programId, mediaId string) error {
	player := GetPlayerClient(playerId)
	if player == nil {
		return fmt.Errorf("player not found")
	}
	return player.PauseLayerMedia(cast.ToUint32(mediaId))
}

func GetProgramMediaProcess(playerId, programId, mediaId string) (int, int, error) {
	player := GetPlayerClient(playerId)
	if player == nil {
		return 0, 0, fmt.Errorf("player not found")
	}
	progress, err := player.QueryLayerProgress(cast.ToUint16(mediaId))
	if err != nil {
		return 0, 0, err
	}
	return int(progress.RemainTime), int(progress.TotalTime), nil
}
func SetProgramMediaProcess(playerId, programId, mediaId string, remainTime, totalTime int) error {
	player := GetPlayerClient(playerId)
	if player == nil {
		return fmt.Errorf("player not found")
	}
	return player.ControlLayerProgress(programId, uint32(remainTime), uint32(totalTime), cast.ToUint16(mediaId))
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
