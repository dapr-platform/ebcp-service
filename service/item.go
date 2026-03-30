package service

import (
	"context"
	"ebcp-service/model"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
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

func StopService() {
	if stopRefresh != nil {
		close(stopRefresh)
	}
}

func refreshItemStatus(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("refreshItemStatus panic: %v", err)
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

	// 静态展项的状态由手动操作（start/stop API）控制，不通过播放器状态推导
	if item.Type == "static" {
		return nil
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

// startItemCore 启动展项核心逻辑，仅操作展项本身并更新数据库状态，不联动上级
func startItemCore(item *model.Ebcp_exhibition_item) error {
	if item.Type == "static" {
		if err := startStaticItem(item); err != nil {
			return fmt.Errorf("启动静态展项失败: %v", err)
		}
	} else {
		players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(),
			model.Ebcp_playerTableInfo.Name, "item_id="+item.ID)
		if err != nil {
			return fmt.Errorf("获取播放设备信息失败: %v", err)
		}
		if len(players) == 0 {
			return fmt.Errorf("播放设备不存在")
		}
		var errs []string
		for _, player := range players {
			if err := PlayerPlay(&player); err != nil {
				errs = append(errs, fmt.Sprintf("播放设备 %s 播放节目失败: %v", player.ID, err))
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf(strings.Join(errs, "\n"))
		}
	}
	item.Status = ItemStatusStart
	if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
		return err
	}
	if item.Type == "static" {
		syncControlDeviceStatus(item.ID, ItemStatusStart)
	}
	return nil
}

// stopItemCore 停止展项核心逻辑，仅操作展项本身并更新数据库状态，不联动上级
func stopItemCore(item *model.Ebcp_exhibition_item) error {
	if item.Type == "static" {
		if err := stopStaticItem(item); err != nil {
			return fmt.Errorf("停止静态展项失败: %v", err)
		}
	} else {
		players, err := common.DbQuery[model.Ebcp_player](context.Background(), common.GetDaprClient(),
			model.Ebcp_playerTableInfo.Name, "item_id="+item.ID)
		if err != nil {
			return fmt.Errorf("获取播放设备信息失败: %v", err)
		}
		if len(players) == 0 {
			return fmt.Errorf("播放设备不存在")
		}
		var errs []string
		for _, player := range players {
			if err := PlayerStop(&player); err != nil {
				errs = append(errs, fmt.Sprintf("播放设备 %s 停止节目失败: %v", player.ID, err))
			}
		}
		if len(errs) > 0 {
			return fmt.Errorf(strings.Join(errs, "\n"))
		}
	}
	item.Status = ItemStatusStop
	if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
		return err
	}
	if item.Type == "static" {
		syncControlDeviceStatus(item.ID, ItemStatusStop)
	}
	return nil
}

// StartExhibitionItem 启动单个展项（API/调度级别，包含上级状态联动）
func StartExhibitionItem(id string) error {
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, "id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}
	if err := startItemCore(item); err != nil {
		return err
	}
	// 启动时向上无条件传播：展厅和展览都设为启动
	if item.RoomID != "" {
		if err := UpdateRoomStatus(item.RoomID, ItemStatusStart); err != nil {
			common.Logger.Errorf("更新展厅状态失败: %v", err)
		}
	}
	if item.ExhibitionID != "" {
		if err := UpdateExhibitionStatus(item.ExhibitionID, ItemStatusStart); err != nil {
			common.Logger.Errorf("更新展览状态失败: %v", err)
		}
	}
	return nil
}

// StopExhibitionItem 停止单个展项（API/调度级别，包含上级状态联动）
func StopExhibitionItem(id string) error {
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, "id="+id)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}
	if err := stopItemCore(item); err != nil {
		return err
	}
	// 检查该展厅下所有展项是否都已停止，满足条件时同步展厅状态
	if item.RoomID != "" {
		if err := SyncRoomStatusByItems(item.RoomID); err != nil {
			common.Logger.Errorf("同步展厅状态失败: %v", err)
		}
	}
	// 检查展览下所有展厅是否都已停止，满足条件时同步展览状态
	if item.ExhibitionID != "" {
		if err := SyncExhibitionStatusByRooms(item.ExhibitionID); err != nil {
			common.Logger.Errorf("同步展览状态失败: %v", err)
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

// syncControlDeviceStatus 同步静态展项下所有中控设备的状态
func syncControlDeviceStatus(itemID string, status int32) {
	devices, err := common.DbQuery[model.Ebcp_control_device](context.Background(), common.GetDaprClient(),
		model.Ebcp_control_deviceTableInfo.Name, model.Ebcp_control_device_FIELD_NAME_item_id+"="+itemID)
	if err != nil {
		common.Logger.Errorf("查询展项 %s 的中控设备失败: %v", itemID, err)
		return
	}

	var toUpdate []model.Ebcp_control_device
	for _, device := range devices {
		if device.Status != status {
			toUpdate = append(toUpdate, device)
		}
	}
	if len(toUpdate) == 0 {
		return
	}

	workers := maxWorkers
	if len(toUpdate) < workers {
		workers = len(toUpdate)
	}
	jobs := make(chan model.Ebcp_control_device, len(toUpdate))
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range jobs {
				d.Status = status
				d.UpdatedTime = common.LocalTime(time.Now())
				if err := common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(),
					d, model.Ebcp_control_deviceTableInfo.Name, "id"); err != nil {
					common.Logger.Errorf("更新中控设备 %s 状态失败: %v", d.ID, err)
				}
			}
		}()
	}
	for _, d := range toUpdate {
		jobs <- d
	}
	close(jobs)
	wg.Wait()
}

// SyncItemStatusByDevices 根据展项下所有中控设备的状态，条件性同步展项及上级状态
// 所有设备都停止 → 展项=Stop → 继续向上条件性传播（展厅→展览）
func SyncItemStatusByDevices(itemID string) error {
	if itemID == "" {
		return nil
	}
	devices, err := common.DbQuery[model.Ebcp_control_device](context.Background(), common.GetDaprClient(),
		model.Ebcp_control_deviceTableInfo.Name, model.Ebcp_control_device_FIELD_NAME_item_id+"="+itemID)
	if err != nil {
		return fmt.Errorf("查询中控设备失败: %v", err)
	}
	if len(devices) == 0 {
		return nil
	}
	allStopped := true
	for _, d := range devices {
		if d.Status != ItemStatusStop {
			allStopped = false
			break
		}
	}
	if !allStopped {
		return nil
	}
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, "id="+itemID)
	if err != nil || item == nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item.Status != ItemStatusStop {
		item.Status = ItemStatusStop
		item.UpdatedTime = common.LocalTime(time.Now())
		if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
			*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
			return fmt.Errorf("更新展项状态失败: %v", err)
		}
		common.Logger.Infof("展项 [%s] 下所有设备已停止，同步展项状态为停止", itemID)
	}
	if item.RoomID != "" {
		if err := SyncRoomStatusByItems(item.RoomID); err != nil {
			common.Logger.Errorf("同步展厅状态失败: %v", err)
		}
	}
	if item.ExhibitionID != "" {
		if err := SyncExhibitionStatusByRooms(item.ExhibitionID); err != nil {
			common.Logger.Errorf("同步展览状态失败: %v", err)
		}
	}
	return nil
}

// PropagateDeviceStartUpward 中控设备启动后，无条件向上传播：展项→展厅→展览
func PropagateDeviceStartUpward(itemID string) error {
	if itemID == "" {
		return nil
	}
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name, "id="+itemID)
	if err != nil || item == nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item.Status != ItemStatusStart {
		item.Status = ItemStatusStart
		item.UpdatedTime = common.LocalTime(time.Now())
		if err := common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
			*item, model.Ebcp_exhibition_itemTableInfo.Name, "id"); err != nil {
			return fmt.Errorf("更新展项状态失败: %v", err)
		}
		common.Logger.Infof("中控设备启动，无条件设展项 [%s] 为启动", itemID)
	}
	if item.RoomID != "" {
		if err := UpdateRoomStatus(item.RoomID, ItemStatusStart); err != nil {
			common.Logger.Errorf("更新展厅状态失败: %v", err)
		}
	}
	if item.ExhibitionID != "" {
		if err := UpdateExhibitionStatus(item.ExhibitionID, ItemStatusStart); err != nil {
			common.Logger.Errorf("更新展览状态失败: %v", err)
		}
	}
	return nil
}

// PauseExhibitionItem 暂停单个展项
func PauseExhibitionItem(id string) error {
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

	var errs []string
	for _, player := range players {
		if err := PlayerPause(&player); err != nil {
			errs = append(errs, fmt.Sprintf("播放设备 %s 暂停节目失败: %v", player.ID, err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}

	item.Status = ItemStatusPause
	return common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(),
		*item, model.Ebcp_exhibition_itemTableInfo.Name, "id")
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

	workers := 5
	if len(ids) < workers {
		workers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				err := StartExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	var errs []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errs = append(errs, fmt.Sprintf("启动展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
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

	workers := 5
	if len(ids) < workers {
		workers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				err := StopExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	var errs []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errs = append(errs, fmt.Sprintf("停止展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
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

	workers := 5
	if len(ids) < workers {
		workers = len(ids)
	}

	jobs := make(chan string, len(ids))
	results := make(chan result, len(ids))

	for w := 0; w < workers; w++ {
		go func() {
			for id := range jobs {
				err := PauseExhibitionItem(id)
				results <- result{id: id, err: err}
			}
		}()
	}

	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	var errs []string
	for i := 0; i < len(ids); i++ {
		r := <-results
		if r.err != nil {
			errs = append(errs, fmt.Sprintf("暂停展项 %s 失败: %v", r.id, r.err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}
