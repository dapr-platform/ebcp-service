package service

import (
	"context"
	"ebcp-service/model"

	"github.com/dapr-platform/common"
)

func GetDashboardStats() (map[string]int, error) {
	// Get exhibition count
	exhibitionCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_exhibitionTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}

	// Get room count
	roomCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_roomTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}

	// Get item count
	itemCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}

	// Get device count
	deviceCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}
	
	stats := map[string]int{
		"exhibitionCount": int(exhibitionCount),
		"roomCount":       int(roomCount),
		"itemCount":       int(itemCount),
		"deviceCount":     int(deviceCount),
	}

	return stats, nil
}
