package service

import (
	"context"
	"ebcp-service/model"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dapr-platform/common"
)

// 定义工作池大小
const maxWorkers = 10

// 处理单个展项的结果
type itemResult struct {
	itemID string
	err    error
}

func StartHall(itemType string) error {

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"")
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
		common.Logger.Info("没有需要启动的展项")
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStart) < workers {
		workers = len(itemsToStart)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStart))
	results := make(chan itemResult, len(itemsToStart))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StartExhibitionItem(item.ID)
				results <- itemResult{
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
			common.Logger.Infof("展项 [%s] 启动成功", result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分展项启动失败:\n%v", errors)
	}

	return nil
}

func StopHall(itemType string) error {

	if itemType == "1" {
		return stopHallMedia()
	} else if itemType == "2" {
		return stopHallStatic()
	} else {
		err := stopHallMedia()
		if err != nil {
			return err
		}
		err = stopHallStatic()
		if err != nil {
			return err
		}
	}

	return nil
}

func startHallMedia() error {

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"type=media")
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
		common.Logger.Info("没有需要启动的展项")
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStart) < workers {
		workers = len(itemsToStart)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStart))
	results := make(chan itemResult, len(itemsToStart))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StartExhibitionItem(item.ID)
				results <- itemResult{
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
			common.Logger.Infof("展项 [%s] 启动成功", result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分展项启动失败:\n%v", errors)
	}

	return nil

}

func startHallStatic() error {
	devices, err := common.DbQuery[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), model.Ebcp_control_deviceTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询控制设备失败: %v", err)
	}

	for _, device := range devices {
		commands := device.Commands
		if commands == "" {
			continue
		}
		cmds := make([]map[string]string, 0)
		err = json.Unmarshal([]byte(commands), &cmds)
		if err != nil {
			fmt.Println("解析控制设备命令失败: %v", err)
			continue
		}
		for _, cmd := range cmds {
			if cmd["type"] == "start" {
				sendUDPCommand(device.IPAddress, device.Port, cmd["command"])
			}
		}
	}

	return nil

}

func stopHallMedia() error {

	items, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"")
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
		common.Logger.Info("没有需要停止的展项")
		return nil
	}

	// 创建工作池
	workers := maxWorkers
	if len(itemsToStop) < workers {
		workers = len(itemsToStop)
	}

	// 创建任务和结果通道
	jobs := make(chan model.Ebcp_exhibition_item, len(itemsToStop))
	results := make(chan itemResult, len(itemsToStop))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				err := StopExhibitionItem(item.ID)
				results <- itemResult{
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
			common.Logger.Infof("展项 [%s] 停止成功", result.itemID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分展项停止失败:\n%v", errors)
	}

	return nil
}
func stopHallStatic() error {
	devices, err := common.DbQuery[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), model.Ebcp_control_deviceTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询控制设备失败: %v", err)
	}

	for _, device := range devices {
		commands := device.Commands
		if commands == "" {
			continue
		}
		cmds := make([]map[string]string, 0)
		err = json.Unmarshal([]byte(commands), &cmds)
		if err != nil {
			fmt.Println("解析控制设备命令失败: %v", err)
			continue
		}
		for _, cmd := range cmds {
			if cmd["type"] == "stop" {
				sendUDPCommand(device.IPAddress, device.Port, cmd["command"])
			}
		}
	}

	return nil
}
