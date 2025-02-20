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
