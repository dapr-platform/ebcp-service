package service

import (
	"context"
	"ebcp-service/model"
	"fmt"

	"github.com/dapr-platform/common"
	"github.com/spf13/cast"
)

// StartExhibitionItem 启动单个展项
func StartExhibitionItem(id string) error {
	// 获取展项信息
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, "id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "item_id="+id)
	if err != nil {
		return fmt.Errorf("获取播放设备信息失败: %v", err)
	}
	for _, player := range players {
		if player.CurrentProgramID == "" {
			common.Logger.Errorf("播放设备 %s 没有当前节目", player.ID)
			continue
		}
		client := GetPlayerClient(player.ID)
		if client == nil {
			common.Logger.Errorf("播放设备 %s 未找到", player.ID)
			continue
		}

		err = client.PlayProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			common.Logger.Errorf("播放设备 %s 播放节目失败: %v", player.ID, err)
			continue
		}

	}

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
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, "id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "item_id="+id)
	if err != nil {
		return fmt.Errorf("获取播放设备信息失败: %v", err)
	}
	for _, player := range players {
		if player.CurrentProgramID == "" {
			common.Logger.Errorf("播放设备 %s 没有当前节目", player.ID)
			continue
		}
		client := GetPlayerClient(player.ID)
		if client == nil {
			common.Logger.Errorf("播放设备 %s 未找到", player.ID)
			continue
		}

		err = client.StopProgram(cast.ToUint32(player.CurrentProgramID))
		if err != nil {
			common.Logger.Errorf("播放设备 %s 播放节目失败: %v", player.ID, err)
			continue
		}

	}
	// 更新展项状态为停止
	item.Status = 2
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
