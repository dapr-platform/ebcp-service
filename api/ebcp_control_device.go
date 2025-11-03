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

func InitEbcp_control_deviceRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-control-device/page", Ebcp_control_devicePageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-control-device", Ebcp_control_deviceListHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-control-device", UpsertEbcp_control_deviceHandler)

	r.Delete(common.BASE_CONTEXT+"/ebcp-control-device/{id}", DeleteEbcp_control_deviceHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-control-device/batch-delete", batchDeleteEbcp_control_deviceHandler)

	r.Post(common.BASE_CONTEXT+"/ebcp-control-device/batch-upsert", batchUpsertEbcp_control_deviceHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 中控设备
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device/batch-upsert [post]
func batchUpsertEbcp_control_deviceHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Ebcp_control_device
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_control_device")
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

	err = common.DbBatchUpsert[model.Ebcp_control_device](r.Context(), common.GetDaprClient(), entities, model.Ebcp_control_deviceTableInfo.Name, model.Ebcp_control_device_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 中控设备
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param device_type query string false "device_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param room_id query string false "room_id"
// @Param item_id query string false "item_id"
// @Param status query string false "status"
// @Param commands query string false "commands"
// @Param index_num query string false "index_num"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_control_device]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device/page [get]
func Ebcp_control_devicePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_control_device](w, r, common.GetDaprClient(), "o_ebcp_control_device", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 中控设备
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param device_type query string false "device_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param room_id query string false "room_id"
// @Param item_id query string false "item_id"
// @Param status query string false "status"
// @Param commands query string false "commands"
// @Param index_num query string false "index_num"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_control_device} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device [get]
func Ebcp_control_deviceListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_control_device](w, r, common.GetDaprClient(), "o_ebcp_control_device", "id")
}

// @Summary save
// @Description save
// @Tags 中控设备
// @Accept       json
// @Param item body model.Ebcp_control_device true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_control_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device [post]
func UpsertEbcp_control_deviceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ebcp_control_device
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ebcp_control_device")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ebcp_control_device)
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

	err = common.DbUpsert[model.Ebcp_control_device](r.Context(), common.GetDaprClient(), val, model.Ebcp_control_deviceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 中控设备
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ebcp_control_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device/{id} [delete]
func DeleteEbcp_control_deviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ebcp_control_device")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ebcp_control_device", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 中控设备
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device/batch-delete [post]
func batchDeleteEbcp_control_deviceHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ebcp_control_device")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ebcp_control_device", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
