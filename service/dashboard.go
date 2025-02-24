package service

import (
	"context"
	"ebcp-service/model"

	"github.com/dapr-platform/common"
)

func GetDashboardStats() (map[string]int, error) {

	// Get room count
	roomCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_roomTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}

	// Get area count
	areaCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_areaTableInfo.Name, "id","")
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
	cameraCount, err := common.DbGetCount(context.Background(), common.GetDaprClient(), model.Ebcp_cameraTableInfo.Name, "id","")
	if err != nil {
		return nil, err
	}

	stats := map[string]int{
		"roomCount":   int(roomCount),
		"areaCount":   int(areaCount),
		"itemCount":   int(itemCount),
		"deviceCount": int(deviceCount+cameraCount),
	}

	return stats, nil
}
