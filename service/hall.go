package service

import (
	"context"
	"ebcp-service/model"
	"fmt"

	"github.com/dapr-platform/common"
)

// 定义工作池大小（被 room.go, exhibition.go 等文件使用）
const maxWorkers = 10

// StartHall 展馆一键启动，遍历所有展览并调用展览启动
func StartHall(itemType string) error {
	// 查询所有展览
	exhibitions, err := common.DbQuery[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibitionTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展览失败: %v", err)
	}

	if len(exhibitions) == 0 {
		common.Logger.Info("没有展览需要启动")
		return nil
	}

	// 收集错误
	var errors []string
	for _, exhibition := range exhibitions {
		common.Logger.Infof("开始启动展览 [%s]", exhibition.ID)
		err := StartExhibition(exhibition.ID, itemType)
		if err != nil {
			errors = append(errors, fmt.Sprintf("启动展览 [%s] 失败: %v", exhibition.ID, err))
		} else {
			common.Logger.Infof("展览 [%s] 启动成功", exhibition.ID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分展览启动失败:\n%v", errors)
	}

	return nil
}

// StopHall 展馆一键停止，遍历所有展览并调用展览停止
func StopHall(itemType string) error {
	// 查询所有展览
	exhibitions, err := common.DbQuery[model.Ebcp_exhibition](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibitionTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("查询展览失败: %v", err)
	}

	if len(exhibitions) == 0 {
		common.Logger.Info("没有展览需要停止")
		return nil
	}

	// 收集错误
	var errors []string
	for _, exhibition := range exhibitions {
		common.Logger.Infof("开始停止展览 [%s]", exhibition.ID)
		err := StopExhibition(exhibition.ID, itemType)
		if err != nil {
			errors = append(errors, fmt.Sprintf("停止展览 [%s] 失败: %v", exhibition.ID, err))
		} else {
			common.Logger.Infof("展览 [%s] 停止成功", exhibition.ID)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("部分展览停止失败:\n%v", errors)
	}

	return nil
}
