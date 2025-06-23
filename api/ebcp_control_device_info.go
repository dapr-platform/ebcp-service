package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_control_device_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-control-device-info/page", Ebcp_control_device_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-control-device-info", Ebcp_control_device_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 中控设备详细视图，包含设备信息及其关联的展厅、展览、展馆和展项信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param device_type query string false "device_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param status query string false "status"
// @Param commands query string false "commands"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_status query string false "room_status"
// @Param room_remarks query string false "room_remarks"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param room_location query string false "room_location"
// @Param room_location_value query string false "room_location_value"
// @Param room_location_name query string false "room_location_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param exhibition_status query string false "exhibition_status"
// @Param exhibition_hall_id query string false "exhibition_hall_id"
// @Param exhibition_hall_name query string false "exhibition_hall_name"
// @Param exhibition_hall_remarks query string false "exhibition_hall_remarks"
// @Param linked_item query string false "linked_item"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_control_device_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device-info/page [get]
func Ebcp_control_device_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_control_device_info](w, r, common.GetDaprClient(), "v_ebcp_control_device_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 中控设备详细视图，包含设备信息及其关联的展厅、展览、展馆和展项信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param device_type query string false "device_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param status query string false "status"
// @Param commands query string false "commands"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_status query string false "room_status"
// @Param room_remarks query string false "room_remarks"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param room_location query string false "room_location"
// @Param room_location_value query string false "room_location_value"
// @Param room_location_name query string false "room_location_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param exhibition_status query string false "exhibition_status"
// @Param exhibition_hall_id query string false "exhibition_hall_id"
// @Param exhibition_hall_name query string false "exhibition_hall_name"
// @Param exhibition_hall_remarks query string false "exhibition_hall_remarks"
// @Param linked_item query string false "linked_item"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_control_device_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-control-device-info [get]
func Ebcp_control_device_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_control_device_info](w, r, common.GetDaprClient(), "v_ebcp_control_device_info", "id")
}
