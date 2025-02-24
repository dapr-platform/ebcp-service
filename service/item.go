package service

import (
	"context"
	"ebcp-service/model"
	"fmt"

	"github.com/dapr-platform/common"
)

// StartExhibitionItem 启动单个展项
func StartExhibitionItem(id string) error {
	// 获取展项信息
	item,err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, id+"="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}

	//TODO

	// 更新展项状态为启动
	item.Status = 1
	err = common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), *item, model.Ebcp_exhibition_itemTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新展项状态失败: %v", err)
	}

	return nil
}

// StopExhibitionItem 停止单个展项
func StopExhibitionItem(id string) error {
	// 获取展项信息
	item,err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, id+"="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	//TODO
	// 更新展项状态为停止
	item.Status = 0
	err = common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), *item, model.Ebcp_exhibition_itemTableInfo.Name, "id")
	if err != nil {
		return fmt.Errorf("更新展项状态失败: %v", err)
	}

	return nil
}

// BatchStartExhibitionItems 批量启动展项
func BatchStartExhibitionItems(ids []string) error {
	for _, id := range ids {
		err := StartExhibitionItem(id)
		if err != nil {
			return fmt.Errorf("启动展项 %s 失败: %v", id, err)
		}
	}
	return nil
}

// BatchStopExhibitionItems 批量停止展项
func BatchStopExhibitionItems(ids []string) error {
	for _, id := range ids {
		err := StopExhibitionItem(id)
		if err != nil {
			return fmt.Errorf("停止展项 %s 失败: %v", id, err)
		}
	}
	return nil
}
