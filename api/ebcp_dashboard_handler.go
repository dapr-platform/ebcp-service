package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcpDashboardRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/ebcp-dashboard/stats", GetDashboardStatsHandler)
}

// @Summary Get dashboard statistics
// @Description Get counts of rooms, areas, items and devices
// @Tags 数据统计
// @Produce json
// @Success 200 {object} common.Response{data=map[string]int} "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-dashboard/stats [get]
func GetDashboardStatsHandler(w http.ResponseWriter, r *http.Request) {

	stats, err := service.GetDashboardStats()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(stats))
}
