package api

import (
	"ebcp-service/model"
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)
func InitEbcp_schedule_job_extRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp_schedule_job/{rel-id}/batch-save", batchSaveEbcp_schedule_jobHandler)
}

// @Summary 批量保存定时任务
// @Description 批量保存定时任务，删除原有任务，新增新的任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Param rel-id path string true "rel-id"
// @Param ebcp_schedule_jobs body []model.Ebcp_schedule_job true "ebcp_schedule_job"
// @Success 200 {string} common.Response ""
// @Router /ebcp_schedule_job/{rel-id}/batch-save [post]
func batchSaveEbcp_schedule_jobHandler(w http.ResponseWriter, r *http.Request) {

	relId := chi.URLParam(r, "rel-id")
	if relId == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("item-id is required"))
		return
	}

	var ebcp_schedule_jobs []model.Ebcp_schedule_job
	err := common.ReadRequestBody(r, &ebcp_schedule_jobs)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.BatchSaveEbcp_schedule_job(relId, ebcp_schedule_jobs)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK)
}