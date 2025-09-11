package service

import (
	"context"
	"ebcp-service/model"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dapr-platform/common"
)

const (
	ItemStatusStart = 0
	ItemStatusPause = 1
	ItemStatusStop  = 2
)

var (
	stopRefresh chan struct{}
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("item init panic: %v", err)
		}
	}()
	stopRefresh = make(chan struct{})
	go refreshItemStatus(context.Background())
}

// StopService 停止服务
func StopService() {
	if stopRefresh != nil {
		close(stopRefresh)
	}
}

func refreshItemStatus(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("refreshItemStatus panic: %v", err)
			// 重启该goroutine
			go refreshItemStatus(context.Background())
		}
	}()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-stopRefresh:
			return
		case <-ticker.C:
			if err := doRefreshItemStatus(ctx); err != nil {
				common.Logger.Errorf("刷新展项状态失败: %v", err)
			}
		}
	}
}

func doRefreshItemStatus(ctx context.Context) error {
	// 这里只查询status=1的展项,可能会漏掉其他状态的展项
	// 建议去掉status=1的过滤条件,查询所有展项
	items, err := common.DbQuery[model.Ebcp_exhibition_item](ctx, common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, "")
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}

	for i := range items {
		if err := updateItemStatus(ctx, &items[i]); err != nil {
			common.Logger.Errorf("更新展项 %s 状态失败: %v", items[i].ID, err)
			continue
		}
	}
	return nil
}

func updateItemStatus(ctx context.Context, item *model.Ebcp_exhibition_item) error {
	if item == nil {
		return fmt.Errorf("展项信息为空")
	}

	players, err := common.DbQuery[model.Ebcp_player](ctx, common.GetDaprClient(), model.Ebcp_playerTableInfo.Name, "item_id="+item.ID)
	if err != nil {
		return fmt.Errorf("获取播放设备信息失败: %v", err)
	}

	var newStatus int32 = ItemStatusStop
	for _, player := range players {
		if player.Status == PlayerStatusOnline {
			if player.CurrentProgramState == ProgramStatePlay {
				newStatus = ItemStatusStart
				break
			}
			if player.CurrentProgramState == ProgramStatePause {
				newStatus = ItemStatusPause
				break
			}
		}
	}
	if item.Status == newStatus {
		return nil
	}
	item.Status = newStatus
	return common.DbUpsert[model.Ebcp_exhibition_item](ctx, common.GetDaprClient(), *item, model.Ebcp_exhibition_itemTableInfo.Name, "id")
}

// StartExhibitionItem 启动单个展项
func StartExhibitionItem(id string) error {
	// 获取展项信息
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}
	if item.Type == "static" {
		err = startStaticItem(item)
		if err != nil {
			return fmt.Errorf("启动静态展项失败: %v", err)
		}
	} else {

		players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(),
			model.Ebcp_playerTableInfo.Name,
			"item_id="+id)
		if err != nil {
			return fmt.Errorf("获取播放设备信息失败: %v", err)
		}
		if len(players) == 0 {
			return fmt.Errorf("播放设备不存在")
		}

		var errors []string
		for _, player := range players {
			if err := PlayerPlay(&player); err != nil {
				msg := fmt.Sprintf("播放设备 %s 播放节目失败: %v", player.ID, err)
				common.Logger.Error(msg)
				errors = append(errors, msg)
				continue
			}
		}

		if len(errors) > 0 {
			return fmt.Errorf(strings.Join(errors, "\n"))
		}

		// 更新展项状态为启动
		item.Status = ItemStatusStart
		if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
			*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
			return fmt.Errorf("更新展项状态失败: %v", err)
		}
	}
	roomid := item.RoomID
	if roomid != "" {
		if err := UpdateRoomStatus(roomid, ItemStatusStart); err != nil {
			return fmt.Errorf("更新展室状态失败: %v", err)
		}
	}
	exhibitionid := item.ExhibitionID
	if exhibitionid != "" {
		if err := UpdateExhibitionStatus(exhibitionid, ItemStatusStart); err != nil {
			return fmt.Errorf("更新展览状态失败: %v", err)
		}
	}

	return nil
}
func startStaticItem(item *model.Ebcp_exhibition_item) error {
	commands := item.Commands
	if commands == "" {
		return fmt.Errorf("展项命令为空")
	}
	ip_address := item.IPAddress
	port := item.Port
	if ip_address == "" || port == 0 {
		return fmt.Errorf("展项IP地址或端口为空")
	}
	var commandList []map[string]string
	err := json.Unmarshal([]byte(commands), &commandList)
	if err != nil {
		return fmt.Errorf("解析展项命令失败: %v", err)
	}
	for _, command := range commandList {
		if command["type"] == "start" {
			SendUDPCommand(ip_address, port, command["command"])
		}
	}
	return nil
}
func stopStaticItem(item *model.Ebcp_exhibition_item) error {
	commands := item.Commands
	if commands == "" {
		return fmt.Errorf("展项命令为空")
	}
	ip_address := item.IPAddress
	port := item.Port
	if ip_address == "" || port == 0 {
		return fmt.Errorf("展项IP地址或端口为空")
	}
	var commandList []map[string]string
	err := json.Unmarshal([]byte(commands), &commandList)
	if err != nil {
		return fmt.Errorf("解析展项命令失败: %v", err)
	}
	for _, command := range commandList {
		if command["type"] == "stop" {
			SendUDPCommand(ip_address, port, command["command"])
		}
	}
	return nil
}
func pauseStaticItem(item *model.Ebcp_exhibition_item) error {
	return nil
}

// StopExhibitionItem 停止单个展项
func StopExhibitionItem(id string) error {
	// 获取展项信息
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}
	if item.Type == "static" {
		err = stopStaticItem(item)
		if err != nil {
			return fmt.Errorf("停止静态展项失败: %v", err)
		}
	} else {
		players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(),
			model.Ebcp_playerTableInfo.Name,
			"item_id="+id)
		if err != nil {
			return fmt.Errorf("获取播放设备信息失败: %v", err)
		}
		if len(players) == 0 {
			return fmt.Errorf("播放设备不存在")
		}

		var errors []string
		for _, player := range players {
			if err := PlayerStop(&player); err != nil {
				msg := fmt.Sprintf("播放设备 %s 停止节目失败: %v", player.ID, err)
				common.Logger.Error(msg)
				errors = append(errors, msg)
				continue
			}
		}

		if len(errors) > 0 {
			return fmt.Errorf(strings.Join(errors, "\n"))
		}

		// 更新展项状态为停止
		item.Status = ItemStatusStop
		if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
			*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
			return fmt.Errorf("更新展项状态失败: %v", err)
		}
	}
	return UpdateItemRoomExhibitionStopStatus(item, nil, nil)

}

// PauseExhibitionItem 暂停单个展项
func PauseExhibitionItem(id string) error {
	// 获取展项信息
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}
	if item.Type == "static" {
		return pauseStaticItem(item)
	}

	players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(),
		model.Ebcp_playerTableInfo.Name,
		"item_id="+id)
	if err != nil {
		return fmt.Errorf("获取播放设备信息失败: %v", err)
	}
	if len(players) == 0 {
		return fmt.Errorf("播放设备不存在")
	}

	var errors []string
	for _, player := range players {
		if err := PlayerPause(&player); err != nil {
			msg := fmt.Sprintf("播放设备 %s 暂停节目失败: %v", player.ID, err)
			common.Logger.Error(msg)
			errors = append(errors, msg)
			continue
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	// 更新展项状态为暂停
	item.Status = ItemStatusPause
	if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
		return fmt.Errorf("更新展项状态失败: %v", err)
	}

	return nil
}

// BatchStartExhibitionItems 批量启动展项
func BatchStartExhibitionItems(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("展项ID列表为空")
	}

	type result struct {
		id  string
		err error
	}

	// 使用worker pool限制并发数
	maxWorkers := 5
	if len(ids) < maxWorkers {
		maxWorkers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	// 启动workers
	for w := 0; w < maxWorkers; w++ {
		go func() {
			for id := range jobs {
				err := StartExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	// 发送任务
	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	// 收集结果
	var errors []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errors = append(errors, fmt.Sprintf("启动展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

// BatchStopExhibitionItems 批量停止展项
func BatchStopExhibitionItems(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("展项ID列表为空")
	}

	type result struct {
		id  string
		err error
	}

	// 使用worker pool限制并发数
	maxWorkers := 5
	if len(ids) < maxWorkers {
		maxWorkers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	// 启动workers
	for w := 0; w < maxWorkers; w++ {
		go func() {
			for id := range jobs {
				err := StopExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	// 发送任务
	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	// 收集结果
	var errors []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errors = append(errors, fmt.Sprintf("停止展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

// BatchPauseExhibitionItems 批量暂停展项
func BatchPauseExhibitionItems(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("展项ID列表为空")
	}

	type result struct {
		id  string
		err error
	}

	// 使用worker pool限制并发数
	maxWorkers := 5
	if len(ids) < maxWorkers {
		maxWorkers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	// 启动workers
	for w := 0; w < maxWorkers; w++ {
		go func() {
			for id := range jobs {
				err := PauseExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	// 发送任务
	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	// 收集结果
	var errors []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errors = append(errors, fmt.Sprintf("暂停展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

func UpdateItemRoomExhibitionStopStatus(item *model.Ebcp_exhibition_item, room *model.Ebcp_exhibition_room, exhibition *model.Ebcp_exhibition) error {
	if item != nil {
		roomItems, err := common.DbQuery[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
			model.Ebcp_exhibition_itemTableInfo.Name,
			"room_id="+item.RoomID)
		if err != nil {
			return fmt.Errorf("获取展项信息失败: %v", err)
		}
		status := ItemStatusStop
		for _, roomItem := range roomItems {
			if roomItem.Status != ItemStatusStop {
				status = ItemStatusStart
				break
			}
		}
		if status == ItemStatusStop {
			queryRoom, err := common.DbGetOne[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
				model.Ebcp_exhibition_roomTableInfo.Name,
				"id="+item.RoomID)
			room = queryRoom
			if err != nil {
				return fmt.Errorf("获取展室信息失败: %v", err)
			}
			if queryRoom == nil {
				return fmt.Errorf("展室不存在")
			}
			if err := UpdateRoomStatus(item.RoomID, ItemStatusStop); err != nil {
				return fmt.Errorf("更新展室状态失败: %v", err)
			}
		}
	}
	if room != nil {
		exhibitionRooms, err := common.DbQuery[model.Ebcp_exhibition_room](context.Background(), common.GetDaprClient(),
			model.Ebcp_exhibition_roomTableInfo.Name,
			"exhibition_id="+room.ExhibitionID)
		if err != nil {
			return fmt.Errorf("获取所有展室信息失败: %v", err)
		}
		exhibitionStatus := ItemStatusStop
		for _, exhibitionRoom := range exhibitionRooms {
			if exhibitionRoom.Status != ItemStatusStop {
				exhibitionStatus = ItemStatusStart
				break
			}
		}
		if exhibitionStatus == ItemStatusStop {
			err := UpdateExhibitionStatus(room.ExhibitionID, ItemStatusStop)
			if err != nil {
				return fmt.Errorf("更新展览状态失败: %v", err)
			}
		}
	}
	if exhibition != nil {
		err := UpdateExhibitionStatus(exhibition.ID, ItemStatusStop)
		if err != nil {
			return fmt.Errorf("更新展览状态失败: %v", err)
		}
	}

	return nil
}
