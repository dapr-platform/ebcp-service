package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitDebugRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/debug/schedule-day-judge", DebugScheduleDayJudgeHandler)
	r.Post(common.BASE_CONTEXT+"/debug/send-udp-command", SendUDPCommandHandler)
}

type SendUDPCommandRequest struct {
	IP      string `json:"ip"`
	Port    int32  `json:"port"`
	Command string `json:"command"`
}

// @Summary 发送UDP命令
// @Description 发送UDP命令
// @Tags 调试接口
// @Accept json
// @Produce json
// @Param request body SendUDPCommandRequest true "请求参数"
// @Success 200 {string} string "ok"
// @Router /debug/send-udp-command [post]
func SendUDPCommandHandler(w http.ResponseWriter, r *http.Request) {
	var request SendUDPCommandRequest
	err := common.ReadRequestBody(r, &request)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.SendUDPCommand(request.IP, request.Port, request.Command)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)
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
