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
		playerMu.RLock()
		for playerId, playerClient := range playerClients {
			go func(id string, client *client.PlayerClient) {
				programs, err := client.GetProgramList()
				if err != nil {
					common.Logger.Errorf("Failed to get program list from player %s: %v", id, err)
					return
				}
				exists, err := common.DbQuery[model.Ebcp_player_program](context.Background(), common.GetDaprClient(),
					model.Ebcp_player_programTableInfo.Name, "player_id = "+ id)
				if err != nil {
					common.Logger.Errorf("Failed to check if player %s has programs: %v", id, err)
					return
				}
				addedMap := make(map[string]bool)

				
				// Insert new programs
				for _, program := range programs.Programs {
					if program.IsEmpty {
						continue
					}
					playerProgram := model.Ebcp_player_program{
						ID:          common.GetMD5Hash(id+"_"+cast.ToString(program.ID)),
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
					}
					addedMap[cast.ToString(program.ID)] = true
				}
				// Delete programs that are no longer in the program list
				for _, program := range exists {
					if !addedMap[cast.ToString(program.ProgramID)] {
						err = common.DbDelete(context.Background(), common.GetDaprClient(),
							model.Ebcp_player_programTableInfo.Name, "id = ?", program.ID)
						if err != nil {
							common.Logger.Errorf("Failed to delete program for player %s: %v", id, err)
						}
					}
				}
			}(playerId, playerClient)
		}
		playerMu.RUnlock()
	}
}

// checkConnections verifies and refreshes player connections
func checkConnections() {
	playerMu.Lock()
	defer playerMu.Unlock()

	// Get all players from database,
	players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "")
	if err != nil {
		common.Logger.Errorf("Failed to query players from database: %v", err)
		return
	}

	// Update player clients
	for _, player := range players {
		id := player.ID
		address := player.IPAddress + ":" + cast.ToString(player.Port)

		if _, exists := playerClients[id]; !exists {
			// Create new client
			client, err := client.NewTCPClient(address)
			if err != nil {
				common.Logger.Errorf("Failed to connect to player %s at %s: %v", id, address, err)
				continue
			}
			playerClients[id] = client
		}
	}

	// Remove disconnected clients
	for id := range playerClients {
		found := false
		for _, player := range players {
			if player.ID == id {
				found = true
				break
			}
		}
		if !found {
			playerClients[id].Close()
			delete(playerClients, id)
		}
	}
}
