package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"sync"
	"time"

	"github.com/dapr-platform/common"
)

// 处理单个展项的结果
type roomItemResult struct {
	itemID string
	err    error
}

func StartExhibitionRoom(roomID string, itemType string) error {
	query := "room_id=" + roomID
	if itemType != "" {
		query = query + "&type=" + itemType
	}
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name,
		"id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展室信息失败: %v", err)
	}
	if room == nil {
		return fmt.Errorf("展室不存在")
	}
	if err := UpdateRoomStatus(roomID, ItemStatusStart); err != nil {
		return fmt.Errorf("更新展室状态失败: %v", err)
	}
	if err := UpdateExhibitionStatus(room.ExhibitionID, ItemStatusStart); err != nil {
		return fmt.Errorf("更新展览状态失败: %v", err)
	}
	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		query)
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

	if len(itemsToStart) == 0 {
		common.Logger.Infof("展室 %s 没有需要启动的展项", roomID)
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStart) < workers {
		workers = len(itemsToStart)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStart))
	results := make(chan roomItemResult, len(itemsToStart))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StartExhibitionItem(item.ID)
				results <- roomItemResult{
					itemID: item.ID,
					err:    err,
				}
			}
		}()
	}

	// 发送任务
	for _, item := range itemsToStart {
		jobs <- item
	}
	close(jobs)

	// 等待所有工作协程完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集错误
	var errors []string
	for result := range results {
		if result.err != nil {
			errors = append(errors, fmt.Sprintf("启动展项 [%s] 失败: %v", result.itemID, result.err))
		} else {
			common.Logger.Infof("展室 [%s] 展项 [%s] 启动成功", roomID, result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("展室 %s 部分展项启动失败:\n%v", roomID, errors)
	}

	return nil
}

func StopExhibitionRoom(roomID string, itemType string) error {
	query := "room_id=" + roomID
	if itemType != "" {
		query = query + "&type=" + itemType
	}
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name,
		"id="+roomID)
	if err != nil {
		return fmt.Errorf("获取展室信息失败: %v", err)
	}
	if room == nil {
		return fmt.Errorf("展室不存在")
	}
	if err := UpdateRoomStatus(roomID, ItemStatusStop); err != nil {
		return fmt.Errorf("更新展室状态失败: %v", err)
	}
	if err := UpdateItemRoomExhibitionStopStatus(nil, room, nil); err != nil {
		return fmt.Errorf("更新展览状态失败: %v", err)
	}
	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		query)
	if err != nil {
		return fmt.Errorf("查询展项失败: %v", err)
	}

	// 过滤出需要停止的展项
	var itemsToStop []model.Ebcp_exhibition_item
	for _, item := range items {
		if item.Status != ItemStatusStop {
			itemsToStop = append(itemsToStop, item)
		}
	}

	if len(itemsToStop) == 0 {
		common.Logger.Infof("展室 %s 没有需要停止的展项", roomID)
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStop) < workers {
		workers = len(itemsToStop)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStop))
	results := make(chan roomItemResult, len(itemsToStop))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StopExhibitionItem(item.ID)
				results <- roomItemResult{
					itemID: item.ID,
					err:    err,
				}
			}
		}()
	}

	// 发送任务
	for _, item := range itemsToStop {
		jobs <- item
	}
	close(jobs)

	// 等待所有工作协程完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集错误
	var errors []string
	for result := range results {
		if result.err != nil {
			errors = append(errors, fmt.Sprintf("停止展项 [%s] 失败: %v", result.itemID, result.err))
		} else {
			common.Logger.Infof("展室 [%s] 展项 [%s] 停止成功", roomID, result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("展室 %s 部分展项停止失败:\n%v", roomID, errors)
	}

	return nil
}
func UpdateRoomStatus(roomID string, status int32) error {
	room, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name,
		"id="+roomID)
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
