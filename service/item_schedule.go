package service

import (
	"context"
	"ebcp-service/config"
	"ebcp-service/model"
	"fmt"
	"time"

	"github.com/dapr-platform/common"
	"github.com/spf13/cast"
)

const (
	CycleTypeWorkday  = 1 // 工作日
	CycleTypeWeekend  = 2 // 周末
	CycleTypeHoliday  = 3 // 节假日
	CycleTypeClosing  = 4 // 闭馆日
	CycleTypeEveryday = 5 // 每天

	HolidayTypeNational          = 1 // 法定节假日
	HolidayTypeAdjustment        = 2 // 调休工作日
	HolidayTypeWeekendAdjustment = 3 // 周末调休
	HolidayTypeClosing           = 4 // 闭馆日
)

var (
	stopSchedule      chan struct{}
	cacheHolidayDates map[int][]model.Ebcp_holiday_date
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("item_schedule init panic: %v", err)
		}
	}()
	stopSchedule = make(chan struct{})
	go scheduleService(context.Background())
	go scheduleRefreshHolidayDates(context.Background())
}

func scheduleRefreshHolidayDates(ctx context.Context) {
	cacheHolidayDates = make(map[int][]model.Ebcp_holiday_date)
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("scheduleRefreshHolidayDates panic: %v", err)
			// 重启该goroutine
			go scheduleRefreshHolidayDates(context.Background())
		}
	}()

	refreshHolidayDates()
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			refreshHolidayDates()
		}
	}
}

func refreshHolidayDates() {

	for {
		year := getCurrentYear()
		holidays, err := getHolidayDates(year)
		if err != nil {
			common.Logger.Errorf("获取节假日日期失败: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		cacheHolidayDates[year] = holidays
		break
	}
}

// 获取节假日日期
func getHolidayDates(year int) ([]model.Ebcp_holiday_date, error) {
	holidays, err := common.DbQuery[model.Ebcp_holiday_date](context.Background(), common.GetDaprClient(),
		model.Ebcp_holiday_dateTableInfo.Name,
		"year="+cast.ToString(year))
	if err != nil {
		return nil, fmt.Errorf("获取节假日日期失败: %v", err)
	}
	return holidays, nil
}

// 获取当前年份
func getCurrentYear() int {
	now := time.Now()
	return now.Year()
}

// StopScheduleService 停止调度服务
func StopScheduleService() {
	if stopSchedule != nil {
		close(stopSchedule)
	}
}

// 调度服务主循环
func scheduleService(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("scheduleService panic: %v", err)
			// 重启该goroutine
			go scheduleService(context.Background())
		}
	}()
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-stopSchedule:
			return
		case <-ticker.C:
			if err := processSchedules(ctx); err != nil {
				common.Logger.Errorf("处理调度任务失败: %v", err)
			}
		}
	}
}

// 处理所有调度任务
func processSchedules(ctx context.Context) error {
	// 只获取启用的调度任务
	schedules, err := common.DbQuery[model.Ebcp_item_schedule](ctx, common.GetDaprClient(),
		model.Ebcp_item_scheduleTableInfo.Name,
		"")
	if err != nil {
		return fmt.Errorf("获取调度任务失败: %v", err)
	}

	now := time.Now()

	mapSchedule := make(map[string][]model.Ebcp_item_schedule)
	for _, schedule := range schedules {
		mapSchedule[schedule.ItemID] = append(mapSchedule[schedule.ItemID], schedule)
	}

	for itemID, schedules := range mapSchedule {
		if err := processItemSchedule(ctx, itemID, schedules, now); err != nil {
			common.Logger.Errorf("处理调度任务 %s 失败: %v", itemID, err)
			continue
		}
	}
	return nil
}

func processItemSchedule(ctx context.Context, itemID string, schedules []model.Ebcp_item_schedule, now time.Time) error {
	// 首先检查是否有满足时间条件但被禁用的调度计划
	for _, schedule := range schedules {
		// 检查是否满足调度条件（不考虑enabled状态）
		if matchScheduleTime(ctx, &schedule, now) {
			// 如果满足时间条件但是被禁用，则跳过该展项的所有调度计划
			if schedule.Enabled == 0 {
				common.Logger.Infof("展项 %s 存在满足条件但被禁用的调度计划 %s，跳过所有调度", itemID, schedule.ID)
				return nil
			}
		}
	}

	// 如果没有满足条件的禁用计划，则处理所有启用的调度计划
	for _, schedule := range schedules {
		if schedule.Enabled == 1 {
			if err := processSchedule(ctx, &schedule, now); err != nil {
				common.Logger.Errorf("处理调度任务 %s 失败: %v", schedule.ID, err)
				continue
			}
		}
	}
	return nil
}

// matchScheduleTime 检查调度计划是否满足时间条件（不考虑enabled状态）
func matchScheduleTime(ctx context.Context, schedule *model.Ebcp_item_schedule, now time.Time) bool {
	// 首先检查日期类型是否匹配
	if !shouldScheduleWithDateType(schedule, now) {
		return false
	}

	// 解析时间
	startTime, err := parseScheduleTime(schedule.StartTime, now)
	if err != nil {
		common.Logger.Errorf("解析开始时间失败: %v", err)
		return false
	}
	stopTime, err := parseScheduleTime(schedule.StopTime, now)
	if err != nil {
		common.Logger.Errorf("解析停止时间失败: %v", err)
		return false
	}

	// 检查是否在时间范围内
	return isTimeInRange(now, startTime, stopTime)
}

// 处理单个调度任务
func processSchedule(ctx context.Context, schedule *model.Ebcp_item_schedule, now time.Time) error {
	// 检查是否满足调度条件
	if !matchScheduleTime(ctx, schedule, now) {
		return nil
	}

	// 获取展项信息
	item, err := common.DbGetOne[model.Ebcp_exhibition_item](ctx, common.GetDaprClient(),
		model.Ebcp_exhibition_itemTableInfo.Name,
		"id="+schedule.ItemID)
	if err != nil {
		return fmt.Errorf("获取展项信息失败: %v", err)
	}
	if item == nil {
		return fmt.Errorf("展项不存在")
	}

	// 解析时间
	startTime, err := parseScheduleTime(schedule.StartTime, now)
	if err != nil {
		return fmt.Errorf("解析开始时间失败: %v", err)
	}
	stopTime, err := parseScheduleTime(schedule.StopTime, now)
	if err != nil {
		return fmt.Errorf("解析停止时间失败: %v", err)
	}

	// 根据当前时间决定是启动还是停止
	if isTimeInRange(now, startTime, stopTime) {
		if item.Status != ItemStatusStart {
			if err := StartExhibitionItem(schedule.ItemID); err != nil {
				return fmt.Errorf("启动展项失败: %v", err)
			}
			common.Logger.Infof("调度启动展项 %s", schedule.ItemID)
		}
	} else {
		if item.Status != ItemStatusStop {
			if err := StopExhibitionItem(schedule.ItemID); err != nil {
				return fmt.Errorf("停止展项失败: %v", err)
			}
			common.Logger.Infof("调度停止展项 %s", schedule.ItemID)
		}
	}

	return nil
}
// 检查是否满足调度条件
func shouldScheduleWithDateType(schedule *model.Ebcp_item_schedule, now time.Time) bool {
	switch schedule.CycleType {
	case CycleTypeWorkday:
		// 工作日：是普通工作日(周一至周五)或调休工作日，且不是法定节假日
		return (isWorkday(now) || isAdjustmentWorkday(now)) && !isHoliday(now)
	case CycleTypeWeekend:
		// 周末：是周末(周六日)且不是法定节假日
		return isWeekend(now) && !isHoliday(now) && !isAdjustmentWorkday(now)
	case CycleTypeHoliday:
		// 节假日：是法定节假日
		return isHoliday(now)
	case CycleTypeClosing:
		// 闭馆日：是闭馆日且不是法定节假日
		return isCloseDay(now) && !isHoliday(now)
	case CycleTypeEveryday:
		// 每天：无条件满足
		return true
	default:
		// 未知类型：不满足条件
		return false
	}
}

func isHoliday(t time.Time) bool {
	year := t.Year()
	holidays, ok := cacheHolidayDates[year]
	if !ok {
		return false
	}
	for _, holiday := range holidays {
		if time.Time(holiday.Date).Format("2006-01-02") == t.Format("2006-01-02") && holiday.Type == HolidayTypeNational {
			return true
		}
	}
	return false
}

func isCloseDay(t time.Time) bool {

	// 检查是否是周一
	closeWeekDay := config.CLOSE_WEEK_DAY
	if t.Weekday() == time.Weekday(cast.ToInt(closeWeekDay)) {
		return true
	}

	// 检查是否是特殊闭馆日（从数据库配置）
	year := t.Year()
	holidays, ok := cacheHolidayDates[year]
	if !ok {
		return false
	}
	for _, holiday := range holidays {
		if time.Time(holiday.Date).Format("2006-01-02") == t.Format("2006-01-02") && holiday.Type == HolidayTypeClosing {
			return true
		}
	}
	return false
}
func isAdjustmentWorkday(t time.Time) bool {
	year := t.Year()
	holidays, ok := cacheHolidayDates[year]
	if !ok {
		return false
	}
	for _, holiday := range holidays {
		if time.Time(holiday.Date).Format("2006-01-02") == t.Format("2006-01-02") && (holiday.Type == HolidayTypeAdjustment || holiday.Type == HolidayTypeWeekendAdjustment) {
			return true
		}
	}
	return false
}

// 判断是否是工作日（周一到周五,调休日）
func isWorkday(t time.Time) bool {

	weekday := t.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}

// 判断是否是周末（周六和周日）
func isWeekend(t time.Time) bool {
	weekday := t.Weekday()

	return weekday == time.Saturday || weekday == time.Sunday
}

// 解析调度时间（格式：HH:mm）
func parseScheduleTime(timeStr string, now time.Time) (time.Time, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(now.Year(), now.Month(), now.Day(),
		t.Hour(), t.Minute(), 0, 0, now.Location()), nil
}

// 判断当前时间是否在时间范围内
func isTimeInRange(now, start, stop time.Time) bool {
	// 如果停止时间小于开始时间，说明跨天
	if stop.Before(start) {
		stop = stop.AddDate(0, 0, 1)
		if now.Before(start) {
			now = now.AddDate(0, 0, 1)
		}
	}
	return now.After(start) && now.Before(stop)
}
