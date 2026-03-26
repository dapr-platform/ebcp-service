package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"sync"
	"time"

	"github.com/dapr-platform/common"
)

type roomItemResult struct {
	itemID string
	err    error
}

// StartExhibitionRoom 启动展厅：启动所有展项 → 设置展厅为启动 → 条件性同步展览状态
func StartExhibitionRoom(roomID string, itemType string) error {
	query := "room_id=" + roomID
	if itemType != "" {
		query = query + "&type=" + itemType
	}
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展室信息失败: %v", err)
	}
	if room == nil {
		return fmt.Errorf("展室不存在")
	}

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, query)
	if err != nil {
		return fmt.Errorf("查询展项失败: %v", err)
	}

	var itemsToStart []model.Ebcp_exhibition_item
	for _, item := range items {
		if item.Status != ItemStatusStart {
			itemsToStart = append(itemsToStart, item)
		}
	}

	if len(itemsToStart) > 0 {
		startItemsBatch(roomID, itemsToStart, true)
	} else {
		common.Logger.Infof("展室 %s 没有需要启动的展项", roomID)
	}

	if err := UpdateRoomStatus(roomID, ItemStatusStart); err != nil {
		return fmt.Errorf("更新展室状态失败: %v", err)
	}

	// 启动时向上无条件传播：展览也设为启动
	if room.ExhibitionID != "" {
		if err := UpdateExhibitionStatus(room.ExhibitionID, ItemStatusStart); err != nil {
			common.Logger.Errorf("更新展览状态失败: %v", err)
		}
	}

	return nil
}

// StopExhibitionRoom 停止展厅：停止所有展项 → 设置展厅为停止 → 条件性同步展览状态
func StopExhibitionRoom(roomID string, itemType string) error {
	query := "room_id=" + roomID
	if itemType != "" {
		query = query + "&type=" + itemType
	}
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展室信息失败: %v", err)
	}
	if room == nil {
		return fmt.Errorf("展室不存在")
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
		stopItemsBatch(roomID, itemsToStop, true)
	} else {
		common.Logger.Infof("展室 %s 没有需要停止的展项", roomID)
	}

	if err := UpdateRoomStatus(roomID, ItemStatusStop); err != nil {
		return fmt.Errorf("更新展室状态失败: %v", err)
	}

	// 检查展览下所有展厅是否都已停止，满足条件时同步展览状态
	if room.ExhibitionID != "" {
		if err := SyncExhibitionStatusByRooms(room.ExhibitionID); err != nil {
			common.Logger.Errorf("同步展览状态失败: %v", err)
		}
	}

	return nil
}

// startItemsBatch 并发启动一批展项（使用 startItemCore，不触发逐级联动）
func startItemsBatch(contextID string, items []model.Ebcp_exhibition_item, isRoom bool) {
	workers := maxWorkers
	if len(items) < workers {
		workers = len(items)
	}

	jobs := make(chan model.Ebcp_exhibition_item, len(items))
	results := make(chan roomItemResult, len(items))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := startItemCore(&item)
				results <- roomItemResult{itemID: item.ID, err: err}
			}
		}()
	}

	for _, item := range items {
		jobs <- item
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	label := "展室"
	if !isRoom {
		label = "展览"
	}
	for result := range results {
		if result.err != nil {
			common.Logger.Errorf("%s [%s] 启动展项 [%s] 失败: %v", label, contextID, result.itemID, result.err)
		} else {
			common.Logger.Infof("%s [%s] 展项 [%s] 启动成功", label, contextID, result.itemID)
		}
	}
}

// stopItemsBatch 并发停止一批展项（使用 stopItemCore，不触发逐级联动）
func stopItemsBatch(contextID string, items []model.Ebcp_exhibition_item, isRoom bool) {
	workers := maxWorkers
	if len(items) < workers {
		workers = len(items)
	}

	jobs := make(chan model.Ebcp_exhibition_item, len(items))
	results := make(chan roomItemResult, len(items))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := stopItemCore(&item)
				results <- roomItemResult{itemID: item.ID, err: err}
			}
		}()
	}

	for _, item := range items {
		jobs <- item
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	label := "展室"
	if !isRoom {
		label = "展览"
	}
	for result := range results {
		if result.err != nil {
			common.Logger.Errorf("%s [%s] 停止展项 [%s] 失败: %v", label, contextID, result.itemID, result.err)
		} else {
			common.Logger.Infof("%s [%s] 展项 [%s] 停止成功", label, contextID, result.itemID)
		}
	}
}

func UpdateRoomStatus(roomID string, status int32) error {
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展室信息失败: %v", err)
	}
	if room == nil {
		return fmt.Errorf("展室不存在")
	}
	room.Status = status
	room.UpdatedBy = "system"
	room.UpdatedTime = common.LocalTime(time.Now())
	return common.DbUpsert[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		*room, model.Ebcp_exhibition_roomTableInfo.Name, "id")
}

// updateRoomStatusBatch 并发批量更新展厅状态
func updateRoomStatusBatch(rooms []model.Ebcp_exhibition_room, status int32) {
	var toUpdate []model.Ebcp_exhibition_room
	for _, room := range rooms {
		if room.Status != status {
			toUpdate = append(toUpdate, room)
		}
	}
	if len(toUpdate) == 0 {
		return
	}
	workers := maxWorkers
	if len(toUpdate) < workers {
		workers = len(toUpdate)
	}
	jobs := make(chan model.Ebcp_exhibition_room, len(toUpdate))
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range jobs {
				if err := UpdateRoomStatus(r.ID, status); err != nil {
					common.Logger.Errorf("批量更新展厅 %s 状态失败: %v", r.ID, err)
				}
			}
		}()
	}
	for _, r := range toUpdate {
		jobs <- r
	}
	close(jobs)
	wg.Wait()
}

// SyncRoomStatusByItems 根据展厅下所有展项的状态同步展厅状态
// 所有展项都启动 → 展厅=Start；所有展项都停止 → 展厅=Stop；否则不变
func SyncRoomStatusByItems(roomID string) error {
	if roomID == "" {
		return nil
	}
	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, "room_id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展项列表失败: %v", err)
	}
	if len(items) == 0 {
		return nil
	}
	allStarted := true
	allStopped := true
	for _, item := range items {
		if item.Status != ItemStatusStart {
			allStarted = false
		}
		if item.Status != ItemStatusStop {
			allStopped = false
		}
	}
	if allStarted {
		common.Logger.Infof("展厅 %s 下所有展项已启动，同步展厅状态为启动", roomID)
		return UpdateRoomStatus(roomID, ItemStatusStart)
	}
	if allStopped {
		common.Logger.Infof("展厅 %s 下所有展项已停止，同步展厅状态为停止", roomID)
		return UpdateRoomStatus(roomID, ItemStatusStop)
	}
	return nil
}
