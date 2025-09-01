package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dapr-platform/common"
)

const (
	RelTypeExhibition = "exhibition" // 展览类型
	RelTypeRoom       = "room"       // 展厅类型
)

var (
	stopScheduleJob chan struct{}
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("schedule_job init panic: %v", err)
		}
	}()
	stopScheduleJob = make(chan struct{})
	go scheduleJobService(context.Background())

}

func BatchSaveEbcp_schedule_job(relId string, ebcp_schedule_jobs []model.Ebcp_schedule_job) error {
	err := common.DbDeleteByOps(context.Background(), common.GetDaprClient(),
		model.Ebcp_schedule_jobTableInfo.Name,
		[]string{"rel_id"},
		[]string{"=="},
		[]any{relId})
	if err != nil {
		return fmt.Errorf("删除定时任务失败: %v", err)
	}
	if len(ebcp_schedule_jobs) > 0 {
		for i := range ebcp_schedule_jobs {
			if ebcp_schedule_jobs[i].ID == "" {
				ebcp_schedule_jobs[i].ID = common.NanoId()
			}
			// 确保item_id被正确设置
			ebcp_schedule_jobs[i].RelID = relId
		}
		err = common.DbBatchInsert[model.Ebcp_schedule_job](context.Background(), common.GetDaprClient(),
			ebcp_schedule_jobs, model.Ebcp_schedule_jobTableInfo.Name)
		if err != nil {
			return fmt.Errorf("保存定时任务失败: %v", err)
		}
	}

	return nil
}

// 定时任务调度服务主循环
func scheduleJobService(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("scheduleJobService panic: %v", err)
			// 重启该goroutine
			go scheduleJobService(context.Background())
		}
	}()
	common.Logger.Infof("scheduleJobService start")
	// 每分钟检查一次
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-stopScheduleJob:
			return
		case <-ticker.C:
			if err := processScheduleJobs(ctx); err != nil {
				common.Logger.Errorf("处理定时任务失败: %v", err)
			}
		}
	}
}

// 处理所有定时任务
func processScheduleJobs(ctx context.Context) error {
	// 只获取启用的定时任务
	jobs, err := common.DbQuery[model.Ebcp_schedule_job](ctx, common.GetDaprClient(),
		model.Ebcp_schedule_jobTableInfo.Name,
		"enabled=1")
	if err != nil {
		return fmt.Errorf("获取定时任务失败: %v", err)
	}

	now := time.Now()

	for _, job := range jobs {
		if err := processScheduleJob(ctx, &job, now); err != nil {
			common.Logger.Errorf("处理定时任务 %s 失败: %v", job.ID, err)
			continue
		}
	}

	return nil
}

// 处理单个定时任务
func processScheduleJob(ctx context.Context, job *model.Ebcp_schedule_job, now time.Time) error {
	// 检查是否满足执行条件
	shouldExecute, action, err := shouldExecuteJob(job, now)
	if err != nil {
		return fmt.Errorf("检查任务执行条件失败: %v", err)
	}

	if !shouldExecute {
		return nil
	}

	// 根据关联类型和操作执行相应的动作
	switch job.RelType {
	case RelTypeExhibition:
		return executeExhibitionAction(job.RelID, action)
	case RelTypeRoom:
		return executeRoomAction(job.RelID, action)
	default:
		return fmt.Errorf("不支持的关联类型: %s", job.RelType)
	}
}

// 检查是否应该执行任务
func shouldExecuteJob(job *model.Ebcp_schedule_job, now time.Time) (bool, string, error) {
	// 检查日期范围
	if !isInDateRange(job, now) {
		return false, "", nil
	}

	// 检查是否在指定的星期几
	if !isInWeekDays(job, now) {
		return false, "", nil
	}

	// 解析开始和结束时间
	startTime, err := parseTimeString(job.StartTime, now)
	if err != nil {
		return false, "", fmt.Errorf("解析开始时间失败: %v", err)
	}

	stopTime, err := parseTimeString(job.StopTime, now)
	if err != nil {
		return false, "", fmt.Errorf("解析结束时间失败: %v", err)
	}

	// 判断当前时间应该执行的动作
	currentMinute := now.Hour()*60 + now.Minute()
	startMinute := startTime.Hour()*60 + startTime.Minute()
	stopMinute := stopTime.Hour()*60 + stopTime.Minute()

	// 如果当前时间刚好是开始时间（误差1分钟内）
	if abs(currentMinute-startMinute) <= 1 {
		return true, "start", nil
	}

	// 如果当前时间刚好是结束时间（误差1分钟内）
	if abs(currentMinute-stopMinute) <= 1 {
		return true, "stop", nil
	}

	return false, "", nil
}

// 检查是否在日期范围内
func isInDateRange(job *model.Ebcp_schedule_job, now time.Time) bool {
	nowDate := now.Format("2006-01-02")

	// 如果设置了开始日期，检查是否已到开始日期
	if job.StartDate != "" && strings.Compare(nowDate, job.StartDate) < 0 {
		return false
	}

	// 如果设置了结束日期，检查是否已超过结束日期
	if job.StopDate != "" && strings.Compare(nowDate, job.StopDate) > 0 {
		return false
	}

	return true
}

// 检查是否在指定的星期几
func isInWeekDays(job *model.Ebcp_schedule_job, now time.Time) bool {
	if job.WeekDays == "" {
		return true // 如果没有指定星期几，则每天都执行
	}

	// weekdays 格式为逗号分隔的数字，1-7代表周一到周日
	weekdayStrs := strings.Split(job.WeekDays, ",")
	currentWeekday := int(now.Weekday())

	// Go的Weekday: Sunday=0, Monday=1, ..., Saturday=6
	// 数据库格式: Monday=1, Tuesday=2, ..., Sunday=7
	// 转换：Go的0(Sunday) -> 7, Go的1-6 -> 1-6
	if currentWeekday == 0 {
		currentWeekday = 7
	}

	for _, weekdayStr := range weekdayStrs {
		weekdayStr = strings.TrimSpace(weekdayStr)
		if weekdayStr == "" {
			continue
		}

		weekday, err := strconv.Atoi(weekdayStr)
		if err != nil {
			common.Logger.Warnf("解析星期几失败: %s", weekdayStr)
			continue
		}

		if weekday == currentWeekday {
			return true
		}
	}

	return false
}

// 解析时间字符串（格式：HH:mm）
func parseTimeString(timeStr string, now time.Time) (time.Time, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(now.Year(), now.Month(), now.Day(),
		t.Hour(), t.Minute(), 0, 0, now.Location()), nil
}

// 执行展览相关操作
func executeExhibitionAction(exhibitionID, action string) error {
	switch action {
	case "start":
		common.Logger.Infof("定时任务启动展览: %s", exhibitionID)
		return StartExhibition(exhibitionID, "")
	case "stop":
		common.Logger.Infof("定时任务停止展览: %s", exhibitionID)
		return StopExhibition(exhibitionID, "")
	default:
		return fmt.Errorf("不支持的展览操作: %s", action)
	}
}

// 执行展厅相关操作
func executeRoomAction(roomID, action string) error {
	switch action {
	case "start":
		common.Logger.Infof("定时任务启动展厅: %s", roomID)
		return StartExhibitionRoom(roomID, "")
	case "stop":
		common.Logger.Infof("定时任务停止展厅: %s", roomID)
		return StopExhibitionRoom(roomID, "")
	default:
		return fmt.Errorf("不支持的展厅操作: %s", action)
	}
}

// 计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 停止定时任务服务
func StopScheduleJobService() {
	close(stopScheduleJob)
}
