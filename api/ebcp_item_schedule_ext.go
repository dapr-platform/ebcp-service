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

// @Summary 批量保存展项定时任务
// @Description 批量保存展项定时任务，删除原有任务，新增新的任务
// @Tags 展项定时任务
// @Accept json
// @Produce json
// @Param item-id path string true "item-id"
// @Param ebcp_item_schedules body []model.Ebcp_item_schedule true "ebcp_item_schedule"
// @Success 200 {string} common.Response ""
// @Router /ebcp_item_schedule/{item-id}/batch-save [post]
func batchSaveEbcp_item_scheduleHandler(w http.ResponseWriter, r *http.Request) {

	itemId := chi.URLParam(r, "item-id")
	if itemId == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("item-id is required"))
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
