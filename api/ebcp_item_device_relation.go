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

func InitEbcp_item_device_relationRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-item-device-relation/page", Ebcp_item_device_relationPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-item-device-relation", Ebcp_item_device_relationListHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-item-device-relation", UpsertEbcp_item_device_relationHandler)

	r.Delete(common.BASE_CONTEXT+"/ebcp-item-device-relation/{id}", DeleteEbcp_item_device_relationHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-item-device-relation/batch-delete", batchDeleteEbcp_item_device_relationHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-item-device-relation/batch-upsert", batchUpsertEbcp_item_device_relationHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 展项关联配置
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation/batch-upsert [post]
func batchUpsertEbcp_item_device_relationHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Ebcp_item_device_relation
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_item_device_relation")
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

	err = common.DbBatchUpsert[model.Ebcp_item_device_relation](r.Context(), common.GetDaprClient(), entities, model.Ebcp_item_device_relationTableInfo.Name, model.Ebcp_item_device_relation_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展项关联配置
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param exhibition_item_id query string false "exhibition_item_id"
// @Param device_type query string false "device_type"
// @Param device_id query string false "device_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_item_device_relation}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation/page [get]
func Ebcp_item_device_relationPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_item_device_relation](w, r, common.GetDaprClient(), "o_ebcp_item_device_relation", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展项关联配置
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param exhibition_item_id query string false "exhibition_item_id"
// @Param device_type query string false "device_type"
// @Param device_id query string false "device_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_item_device_relation} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation [get]
func Ebcp_item_device_relationListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_item_device_relation](w, r, common.GetDaprClient(), "o_ebcp_item_device_relation", "id")
}

// @Summary save
// @Description save
// @Tags 展项关联配置
// @Accept       json
// @Param item body model.Ebcp_item_device_relation true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_item_device_relation} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation [post]
func UpsertEbcp_item_device_relationHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_item_device_relation
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_item_device_relation")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_item_device_relation)
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

	err = common.DbUpsert[model.Ebcp_item_device_relation](r.Context(), common.GetDaprClient(), val, model.Ebcp_item_device_relationTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 展项关联配置
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_item_device_relation} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation/{id} [delete]
func DeleteEbcp_item_device_relationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_item_device_relation")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ebcp_item_device_relation", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 展项关联配置
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-item-device-relation/batch-delete [post]
func batchDeleteEbcp_item_device_relationHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_item_device_relation")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ebcp_item_device_relation", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
