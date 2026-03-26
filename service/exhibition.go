package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"time"

	"github.com/dapr-platform/common"
)

// StartExhibition 启动展览：启动所有展项 → 设置所有展厅为启动 → 设置展览为启动
func StartExhibition(exhibitionId string, itemType string) error {
	query := "exhibition_id=" + exhibitionId
	if itemType != "" {
		query = query + "&type=" + itemType
	}

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, query)
	if err != nil {
		return fmt.Errorf("查询展项失败: %v", err)
	}

	// 过滤出需要启动的展项
	var itemsToStart []model.Ebcp_exhibition_item
	for _, item := range items {
		if item.Status != ItemStatusStart {
			itemsToStart = append(itemsToStart, item)
		}
	}

	if len(itemsToStart) > 0 {
		startItemsBatch(exhibitionId, itemsToStart, false)
	} else {
		common.Logger.Infof("展览 %s 没有需要启动的展项", exhibitionId)
	}

	// 并发设置所有展厅状态为启动
	rooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "exhibition_id="+exhibitionId)
	if err != nil {
		common.Logger.Errorf("查询展厅失败: %v", err)
	} else {
		updateRoomStatusBatch(rooms, ItemStatusStart)
	}

	if err := UpdateExhibitionStatus(exhibitionId, ItemStatusStart); err != nil {
		return fmt.Errorf("更新展览状态失败: %v", err)
	}
	return nil
}

// StopExhibition 停止展览：停止所有展项 → 设置所有展厅为停止 → 设置展览为停止
func StopExhibition(exhibitionId string, itemType string) error {
	query := "exhibition_id=" + exhibitionId
	if itemType != "" {
		query = query + "&type=" + itemType
	}

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, query)
	if err != nil {
		return fmt.Errorf("查询展项失败: %v", err)
	}

	var itemsToStop []model.Ebcp_exhibition_item
	for _, item := range items {
		if item.Status != ItemStatusStop {
			itemsToStop = append(itemsToStop, item)
		}
	}

	if len(itemsToStop) > 0 {
		stopItemsBatch(exhibitionId, itemsToStop, false)
	} else {
		common.Logger.Infof("展览 %s 没有需要停止的展项", exhibitionId)
	}

	// 并发设置所有展厅状态为停止
	rooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "exhibition_id="+exhibitionId)
	if err != nil {
		common.Logger.Errorf("查询展厅失败: %v", err)
	} else {
		updateRoomStatusBatch(rooms, ItemStatusStop)
	}

	if err := UpdateExhibitionStatus(exhibitionId, ItemStatusStop); err != nil {
		return fmt.Errorf("更新展览状态失败: %v", err)
	}
	return nil
}

func UpdateExhibitionStatus(exhibitionId string, status int32) error {
	exhibition, err := common.DbGetOne[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibitionTableInfo.Name, "id="+exhibitionId)
	if err != nil {
		return fmt.Errorf("获取展览信息失败: %v", err)
	}
	if exhibition == nil {
		return fmt.Errorf("展览不存在")
	}
	exhibition.Status = status
	exhibition.UpdatedBy = "system"
	exhibition.UpdatedTime = common.LocalTime(time.Now())
	return common.DbUpsert[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		*exhibition, model.Ebcp_exhibitionTableInfo.Name, "id")
}

// SyncExhibitionStatusByRooms 根据所有展厅的状态条件性同步展览状态
// 所有展厅都启动 → 展览=Start；所有展厅都停止 → 展览=Stop；否则不变
func SyncExhibitionStatusByRooms(exhibitionID string) error {
	if exhibitionID == "" {
		return nil
	}
	rooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "exhibition_id="+exhibitionID)
	if err != nil {
		return fmt.Errorf("获取展厅列表失败: %v", err)
	}
	if len(rooms) == 0 {
		return nil
	}
	allStarted := true
	allStopped := true
	for _, room := range rooms {
		if room.Status != ItemStatusStart {
			allStarted = false
		}
		if room.Status != ItemStatusStop {
			allStopped = false
		}
	}
	if allStarted {
		common.Logger.Infof("展览 %s 下所有展厅已启动，同步展览状态为启动", exhibitionID)
		return UpdateExhibitionStatus(exhibitionID, ItemStatusStart)
	}
	if allStopped {
		common.Logger.Infof("展览 %s 下所有展厅已停止，同步展览状态为停止", exhibitionID)
		return UpdateExhibitionStatus(exhibitionID, ItemStatusStop)
	}
	common.Logger.Infof("展览 %s 下展厅状态不一致，不更新展览状态", exhibitionID)
	return nil
}
