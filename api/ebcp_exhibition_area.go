package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func InitEbcp_exhibition_areaRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area/page", Ebcp_exhibition_areaPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area", Ebcp_exhibition_areaListHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area", UpsertEbcp_exhibition_areaHandler)
	r.Delete(common.BASE_CONTEXT+"/ebcp-exhibition-area/{id}", DeleteEbcp_exhibition_areaHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area/batch-delete", batchDeleteEbcp_exhibition_areaHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-area/batch-upsert", batchUpsertEbcp_exhibition_areaHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area/groupby", Ebcp_exhibition_areaGroupbyHandler)
}

// @Summary GroupBy
// @Description GroupBy, for example,  _select=level, then return  {level_val1:sum1,level_val2:sum2}, _where can input status=0
// @Tags Ebcp_exhibition_area
// @Param _select query string true "_select"
// @Param _where query string false "_where"
// @Produce  json
// @Success 200 {object} common.Response{data=[]map[string]any} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area/groupby [get]
func Ebcp_exhibition_areaGroupbyHandler(w http.ResponseWriter, r *http.Request) {

	common.CommonGroupby(w, r, common.GetDaprClient(), "o_ebcp_exhibition_area")
}

// @Summary batch update
// @Description batch update
// @Tags Ebcp_exhibition_area
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area/batch-upsert [post]
func batchUpsertEbcp_exhibition_areaHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Ebcp_exhibition_areaTableInfo.Name, model.Ebcp_exhibition_area_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Ebcp_exhibition_area
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_exhibition_area}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area/page [get]
func Ebcp_exhibition_areaPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_area](w, r, common.GetDaprClient(), "o_ebcp_exhibition_area", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Ebcp_exhibition_area
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_area} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area [get]
func Ebcp_exhibition_areaListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_area](w, r, common.GetDaprClient(), "o_ebcp_exhibition_area", "id")
}

// @Summary save
// @Description save
// @Tags Ebcp_exhibition_area
// @Accept       json
// @Param item body model.Ebcp_exhibition_area true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_exhibition_area} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area [post]
func UpsertEbcp_exhibition_areaHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_exhibition_area
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_exhibition_area")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_exhibition_area)
	}

	err = common.DbUpsert[model.Ebcp_exhibition_area](r.Context(), common.GetDaprClient(), val, model.Ebcp_exhibition_areaTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags Ebcp_exhibition_area
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_exhibition_area} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area/{id} [delete]
func DeleteEbcp_exhibition_areaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_exhibition_area")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ebcp_exhibition_area", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags Ebcp_exhibition_area
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area/batch-delete [post]
func batchDeleteEbcp_exhibition_areaHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_exhibition_area")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ebcp_exhibition_area", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
