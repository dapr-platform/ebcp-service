package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitEbcp_exhibition_area_detailsRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/page", Ebcp_exhibition_area_detailsPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area-details", Ebcp_exhibition_area_detailsListHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area-details", UpsertEbcp_exhibition_area_detailsHandler)
	r.Delete(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/{id}", DeleteEbcp_exhibition_area_detailsHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/batch-delete", batchDeleteEbcp_exhibition_area_detailsHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/batch-upsert", batchUpsertEbcp_exhibition_area_detailsHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/groupby", Ebcp_exhibition_area_detailsGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Ebcp_exhibition_area_details
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/groupby [get]
func Ebcp_exhibition_area_detailsGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details")
}

// @Summary batch update
// @Description batch update
// @Tags Ebcp_exhibition_area_details
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/batch-upsert [post]
func batchUpsertEbcp_exhibition_area_detailsHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Ebcp_exhibition_area_detailsTableInfo.Name, model.Ebcp_exhibition_area_details_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Ebcp_exhibition_area_details
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param area_name query string false "area_name"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param exhibition_items query string false "exhibition_items"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_exhibition_area_details}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/page [get]
func Ebcp_exhibition_area_detailsPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_area_details](w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Ebcp_exhibition_area_details
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param area_name query string false "area_name"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param exhibition_items query string false "exhibition_items"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_area_details} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details [get]
func Ebcp_exhibition_area_detailsListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_area_details](w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details", "id")
}

// @Summary save
// @Description save
// @Tags Ebcp_exhibition_area_details
// @Accept       json
// @Param item body model.Ebcp_exhibition_area_details true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_exhibition_area_details} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details [post]
func UpsertEbcp_exhibition_area_detailsHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_exhibition_area_details
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_exhibition_area_details")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_exhibition_area_details)
	}

	err = common.DbUpsert[model.Ebcp_exhibition_area_details](r.Context(), common.GetDaprClient(), val, model.Ebcp_exhibition_area_detailsTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Ebcp_exhibition_area_details
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_exhibition_area_details} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/{id} [delete]
func DeleteEbcp_exhibition_area_detailsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_exhibition_area_details")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Ebcp_exhibition_area_details
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/batch-delete [post]
func batchDeleteEbcp_exhibition_area_detailsHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_exhibition_area_details")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_ebcp_exhibition_area_details", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
