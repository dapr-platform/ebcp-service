package service

import (
	"context"
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

	for range ticker.C {
		// 复制当前客户端映射，避免长时间持有锁
		clientsCopy := make(map[string]*client.PlayerClient)
		playerMu.RLock()
		for id, client := range playerClients {
			clientsCopy[id] = client
		}
		playerMu.RUnlock()

		// 使用复制的映射处理程序更新
		for playerId, playerClient := range clientsCopy {
			go func(id string, client *client.PlayerClient) {
				programs, err := client.GetProgramList()
				if err != nil {
					common.Logger.Errorf("Failed to get program list from player %s: %v", id, err)
					return
				}
				exists, err := common.DbQuery[model.Ebcp_player_program](context.Background(), common.GetDaprClient(),
					model.Ebcp_player_programTableInfo.Name, "player_id="+id)
				if err != nil {
					common.Logger.Errorf("Failed to check if player %s has programs: %v", id, err)
					return
				}
				addedMap := make(map[string]bool)

				firstProgramId := ""

				// Insert new programs
				for _, program := range programs.Programs {
					if program.IsEmpty {
						continue
					}
					if firstProgramId == "" {
						firstProgramId = cast.ToString(program.ID)
					}
					playerProgram := model.Ebcp_player_program{
						ID:          common.GetMD5Hash(id + "_" + cast.ToString(program.ID)),
						CreatedBy:   "admin",
						CreatedTime: common.LocalTime(time.Now()),
						UpdatedBy:   "admin",
						UpdatedTime: common.LocalTime(time.Now()),
						ProgramID:   cast.ToString(program.ID),
						Name:        program.Name,
						PlayerID:    id,
					}
					err = common.DbUpsert[model.Ebcp_player_program](context.Background(), common.GetDaprClient(),
						playerProgram, model.Ebcp_player_programTableInfo.Name, "id")
					if err != nil {
						common.Logger.Errorf("Failed to upsert program for player %s: %v", id, err)
					} else {
						syncPlayerProgramMedia(client, id, cast.ToString(playerProgram.ID), program.ID)
						addedMap[cast.ToString(program.ID)] = true
					}
				}
				// Delete programs that are no longer in the program list
				for _, program := range exists {
					if !addedMap[cast.ToString(program.ProgramID)] {
						err = common.DbDelete(context.Background(), common.GetDaprClient(),
							model.Ebcp_player_programTableInfo.Name, "id", program.ID)
						if err != nil {
							common.Logger.Errorf("Failed to delete program for player %s: %v", id, err)
						}
					}
				}

				if firstProgramId != "" {
					player, err := common.DbGetOne[model.Ebcp_player](context.Background(), common.GetDaprClient(),
						model.Ebcp_playerTableInfo.Name, "id="+id)
					if err != nil {
						common.Logger.Errorf("Failed to get player %s: %v", id, err)
						return
					}
					if player == nil {
						common.Logger.Errorf("Player %s not found", id)
						return
					}
					if player.CurrentProgramID == "" {
						player.CurrentProgramID = firstProgramId
						err = common.DbUpsert[model.Ebcp_player](context.Background(), common.GetDaprClient(), *player, model.Ebcp_playerTableInfo.Name, "id")
						if err != nil {
							common.Logger.Errorf("Failed to update player %s current program id: %v", id, err)
						}
					}
				}

			}(playerId, playerClient)
		}
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
