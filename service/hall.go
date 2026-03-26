package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"sync"

	"github.com/dapr-platform/common"
)

// 定义工作池大小（被 room.go, exhibition.go 等文件使用）
const maxWorkers = 10

// StartHall 展馆一键启动：并发启动所有展览，再补漏处理孤立展厅和展项
func StartHall(itemType string) error {
	exhibitions, err := common.DbQuery[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibitionTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展览失败: %v", err)
	}

	var mu sync.Mutex
	var errors []string
	var wg sync.WaitGroup
	for _, exhibition := range exhibitions {
		wg.Add(1)
		go func(exh model.Ebcp_exhibition) {
			defer wg.Done()
			common.Logger.Infof("开始启动展览 [%s]", exh.ID)
			if err := StartExhibition(exh.ID, itemType); err != nil {
				mu.Lock()
				errors = append(errors, fmt.Sprintf("启动展览 [%s] 失败: %v", exh.ID, err))
				mu.Unlock()
			} else {
				common.Logger.Infof("展览 [%s] 启动成功", exh.ID)
			}
		}(exhibition)
	}
	wg.Wait()

	if err := startOrphanedRoomsAndItems(itemType); err != nil {
		errors = append(errors, fmt.Sprintf("启动孤立展厅/展项失败: %v", err))
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分启动失败:\n%v", errors)
	}
	return nil
}

// StopHall 展馆一键停止：并发停止所有展览，再补漏处理孤立展厅和展项
func StopHall(itemType string) error {
	exhibitions, err := common.DbQuery[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibitionTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展览失败: %v", err)
	}

	var mu sync.Mutex
	var errors []string
	var wg sync.WaitGroup
	for _, exhibition := range exhibitions {
		wg.Add(1)
		go func(exh model.Ebcp_exhibition) {
			defer wg.Done()
			common.Logger.Infof("开始停止展览 [%s]", exh.ID)
			if err := StopExhibition(exh.ID, itemType); err != nil {
				mu.Lock()
				errors = append(errors, fmt.Sprintf("停止展览 [%s] 失败: %v", exh.ID, err))
				mu.Unlock()
			} else {
				common.Logger.Infof("展览 [%s] 停止成功", exh.ID)
			}
		}(exhibition)
	}
	wg.Wait()

	if err := stopOrphanedRoomsAndItems(itemType); err != nil {
		errors = append(errors, fmt.Sprintf("停止孤立展厅/展项失败: %v", err))
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分停止失败:\n%v", errors)
	}
	return nil
}

// stopOrphanedRoomsAndItems 停止未被展览覆盖的展厅和展项（孤立数据兜底）
func stopOrphanedRoomsAndItems(itemType string) error {
	// 停止所有仍未停止的展项
	query := ""
	if itemType != "" {
		query = "type=" + itemType
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
		common.Logger.Infof("补漏：还有 %d 个展项未停止，开始批量停止", len(itemsToStop))
		stopItemsBatch("orphaned", itemsToStop, false)
	}

	// 并发停止所有仍未停止的展厅
	rooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展厅失败: %v", err)
	}
	updateRoomStatusBatch(rooms, ItemStatusStop)
	return nil
}

// startOrphanedRoomsAndItems 启动未被展览覆盖的展厅和展项（孤立数据兜底）
func startOrphanedRoomsAndItems(itemType string) error {
	query := ""
	if itemType != "" {
		query = "type=" + itemType
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
		common.Logger.Infof("补漏：还有 %d 个展项未启动，开始批量启动", len(itemsToStart))
		startItemsBatch("orphaned", itemsToStart, false)
	}

	// 并发启动所有仍未启动的展厅
	rooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_roomTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展厅失败: %v", err)
	}
	updateRoomStatusBatch(rooms, ItemStatusStart)
	return nil
}
