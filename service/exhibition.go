package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"sync"

	"github.com/dapr-platform/common"
)

// 处理单个展项的结果
type exhibitionItemResult struct {
	itemID string
	err    error
}

func StartExhibition(exhibitionId string, itemType string) error {
	query := "exhibition_id=" + exhibitionId
	if itemType != "" {
		query = query + "&type=" + itemType
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
		common.Logger.Infof("展览 %s 没有需要启动的展项", exhibitionId)
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStart) < workers {
		workers = len(itemsToStart)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStart))
	results := make(chan exhibitionItemResult, len(itemsToStart))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StartExhibitionItem(item.ID)
				results <- exhibitionItemResult{
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
			common.Logger.Infof("展览 [%s] 展项 [%s] 启动成功", exhibitionId, result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("展览 %s 部分展项启动失败:\n%v", exhibitionId, errors)
	}

	return nil
}

func StopExhibition(exhibitionId string, itemType string) error {
	query := "exhibition_id=" + exhibitionId
	if itemType != "" {
		query = query + "&type=" + itemType
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
		common.Logger.Infof("展览 %s 没有需要停止的展项", exhibitionId)
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStop) < workers {
		workers = len(itemsToStop)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStop))
	results := make(chan exhibitionItemResult, len(itemsToStop))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StopExhibitionItem(item.ID)
				results <- exhibitionItemResult{
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
			common.Logger.Infof("展览 [%s] 展项 [%s] 停止成功", exhibitionId, result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("展览 %s 部分展项停止失败:\n%v", exhibitionId, errors)
	}

	return nil
}
