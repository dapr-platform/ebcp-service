package api

import (
	"ebcp-service/model"
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_item_schedule_extRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp_item_schedule/batch-save", batchSaveEbcp_item_scheduleHandler)
}

// @Summary 批量保存ebcp_item_schedule
// @Description 批量保存ebcp_item_schedule
// @Accept json
// @Produce json
// @Param itemId path string true "itemId"
// @Param ebcp_item_schedules body []model.Ebcp_item_schedule true "ebcp_item_schedule"
// @Success 200 {string} common.Response ""
// @Router /ebcp_item_schedule/{itemId}/batch-save [post]
func batchSaveEbcp_item_scheduleHandler(w http.ResponseWriter, r *http.Request) {

	itemId := chi.URLParam(r, "itemId")
	if itemId == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("itemId is required"))
		return
	}

	var ebcp_item_schedules []model.Ebcp_item_schedule
	err := common.ReadRequestBody(r, &ebcp_item_schedules)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.BatchSaveEbcp_item_schedule(itemId, ebcp_item_schedules)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)
}
