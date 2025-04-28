package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_hallExtRoute(r chi.Router) {

	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-hall/start", Ebcp_exhibition_hallStartHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-hall/stop", Ebcp_exhibition_hallStopHandler)

}

// @Summary 展馆一键启动
// @Description 展馆一键启动
// @Tags 展馆
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-hall/start [post]
func Ebcp_exhibition_hallStartHandler(w http.ResponseWriter, r *http.Request) {
	err := service.StartHall()
	if err != nil {
		common.HttpResult(w, common.OK.AppendMsg(err.Error()))
	}
	common.HttpResult(w, common.OK)
}

// @Summary 展馆一键停止
// @Description 展馆一键停止
// @Tags 展馆
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-hall/stop [post]
func Ebcp_exhibition_hallStopHandler(w http.ResponseWriter, r *http.Request) {
	err := service.StopHall()
	if err != nil {
		common.HttpResult(w, common.OK.AppendMsg(err.Error()))
	}
	common.HttpResult(w, common.OK)
}
