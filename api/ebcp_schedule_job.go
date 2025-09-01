package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"strings"

	"time"
)

var _ = time.Now()

func InitEbcp_schedule_jobRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-schedule-job/page", Ebcp_schedule_jobPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-schedule-job", Ebcp_schedule_jobListHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-job", UpsertEbcp_schedule_jobHandler)

	r.Delete(common.BASE_CONTEXT+"/ebcp-schedule-job/{id}", DeleteEbcp_schedule_jobHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-job/batch-delete", batchDeleteEbcp_schedule_jobHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-job/batch-upsert", batchUpsertEbcp_schedule_jobHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 定时任务
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job/batch-upsert [post]
func batchUpsertEbcp_schedule_jobHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Ebcp_schedule_job
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_schedule_job")
	if exists {
		for _, v := range entities {
			_, err1 := beforeHook(r, v)
			if err1 != nil {
				common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
				return
			}
		}

	}
	for _, v := range entities {
		if v.ID == "" {
			v.ID = common.NanoId()
		}

		if time.Time(v.CreatedTime).IsZero() {
			v.CreatedTime = common.LocalTime(time.Now())
		}

		if time.Time(v.UpdatedTime).IsZero() {
			v.UpdatedTime = common.LocalTime(time.Now())
		}

	}

	err = common.DbBatchUpsert[model.Ebcp_schedule_job](r.Context(), common.GetDaprClient(), entities, model.Ebcp_schedule_jobTableInfo.Name, model.Ebcp_schedule_job_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 定时任务
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param rel_id query string false "rel_id"
// @Param rel_type query string false "rel_type"
// @Param start_time query string false "start_time"
// @Param stop_time query string false "stop_time"
// @Param start_date query string false "start_date"
// @Param stop_date query string false "stop_date"
// @Param week_days query string false "week_days"
// @Param enabled query string false "enabled"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_schedule_job]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job/page [get]
func Ebcp_schedule_jobPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_schedule_job](w, r, common.GetDaprClient(), "o_ebcp_schedule_job", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 定时任务
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param rel_id query string false "rel_id"
// @Param rel_type query string false "rel_type"
// @Param start_time query string false "start_time"
// @Param stop_time query string false "stop_time"
// @Param start_date query string false "start_date"
// @Param stop_date query string false "stop_date"
// @Param week_days query string false "week_days"
// @Param enabled query string false "enabled"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_schedule_job} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job [get]
func Ebcp_schedule_jobListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_schedule_job](w, r, common.GetDaprClient(), "o_ebcp_schedule_job", "id")
}

// @Summary save
// @Description save
// @Tags 定时任务
// @Accept       json
// @Param item body model.Ebcp_schedule_job true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_schedule_job} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job [post]
func UpsertEbcp_schedule_jobHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_schedule_job
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_schedule_job")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_schedule_job)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}

	if time.Time(val.CreatedTime).IsZero() {
		val.CreatedTime = common.LocalTime(time.Now())
	}

	if time.Time(val.UpdatedTime).IsZero() {
		val.UpdatedTime = common.LocalTime(time.Now())
	}

	err = common.DbUpsert[model.Ebcp_schedule_job](r.Context(), common.GetDaprClient(), val, model.Ebcp_schedule_jobTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 定时任务
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_schedule_job} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job/{id} [delete]
func DeleteEbcp_schedule_jobHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_schedule_job")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ebcp_schedule_job", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 定时任务
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-job/batch-delete [post]
func batchDeleteEbcp_schedule_jobHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of ids is 0"))
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_schedule_job")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ebcp_schedule_job", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
