package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitEbcp_schedule_actionRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/ebcp-schedule-action/page", Ebcp_schedule_actionPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-schedule-action", Ebcp_schedule_actionListHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-action", UpsertEbcp_schedule_actionHandler)
	r.Delete(common.BASE_CONTEXT+"/ebcp-schedule-action/{id}", DeleteEbcp_schedule_actionHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-action/batch-delete", batchDeleteEbcp_schedule_actionHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-schedule-action/batch-upsert", batchUpsertEbcp_schedule_actionHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-schedule-action/groupby", Ebcp_schedule_actionGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Ebcp_schedule_action
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action/groupby [get]
func Ebcp_schedule_actionGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_ebcp_schedule_action")
}

// @Summary batch update
// @Description batch update
// @Tags Ebcp_schedule_action
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action/batch-upsert [post]
func batchUpsertEbcp_schedule_actionHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Ebcp_schedule_actionTableInfo.Name, model.Ebcp_schedule_action_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Ebcp_schedule_action
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param action_type query string false "action_type"
// @Param target_id query string false "target_id"
// @Param operation_details query string false "operation_details"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_schedule_action}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action/page [get]
func Ebcp_schedule_actionPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_schedule_action](w, r, common.GetDaprClient(), "o_ebcp_schedule_action", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Ebcp_schedule_action
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param action_type query string false "action_type"
// @Param target_id query string false "target_id"
// @Param operation_details query string false "operation_details"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_schedule_action} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action [get]
func Ebcp_schedule_actionListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_schedule_action](w, r, common.GetDaprClient(), "o_ebcp_schedule_action", "id")
}

// @Summary save
// @Description save
// @Tags Ebcp_schedule_action
// @Accept       json
// @Param item body model.Ebcp_schedule_action true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_schedule_action} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action [post]
func UpsertEbcp_schedule_actionHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_schedule_action
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_schedule_action")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_schedule_action)
	}

	err = common.DbUpsert[model.Ebcp_schedule_action](r.Context(), common.GetDaprClient(), val, model.Ebcp_schedule_actionTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Ebcp_schedule_action
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_schedule_action} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action/{id} [delete]
func DeleteEbcp_schedule_actionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_schedule_action")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ebcp_schedule_action", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Ebcp_schedule_action
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-schedule-action/batch-delete [post]
func batchDeleteEbcp_schedule_actionHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_schedule_action")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ebcp_schedule_action", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
