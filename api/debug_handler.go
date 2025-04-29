package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitDebugRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/debug/schedule-day-judge", DebugScheduleDayJudgeHandler)
}

// @Summary 测试日期类型接口
// @Description 测试日期类型接口
// @Tags 调试接口
// @Accept json
// @Produce json
// @Param date query string true "日期"
// @Success 200 {string} string "ok"
// @Router /debug/schedule-day-judge [get]
func DebugScheduleDayJudgeHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("date is required"))
		return
	}

	result, err := service.JudgeScheduleDay(date)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(result))

}
