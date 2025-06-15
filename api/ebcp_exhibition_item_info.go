package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_item_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-item-info/page", Ebcp_exhibition_item_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-item-info", Ebcp_exhibition_item_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展项详细视图，包含展项信息及其关联的展厅、展览、设备和定时任务信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param remarks query string false "remarks"
// @Param export_info query string false "export_info"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param room_location query string false "room_location"
// @Param room_location_value query string false "room_location_value"
// @Param room_location_name query string false "room_location_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param player_devices query string false "player_devices"
// @Param control_devices query string false "control_devices"
// @Param schedules query string false "schedules"
// @Param commands query string false "commands"
// @Param sub_type query string false "sub_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_exhibition_item_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-item-info/page [get]
func Ebcp_exhibition_item_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_item_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_item_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展项详细视图，包含展项信息及其关联的展厅、展览、设备和定时任务信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param remarks query string false "remarks"
// @Param export_info query string false "export_info"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param room_location query string false "room_location"
// @Param room_location_value query string false "room_location_value"
// @Param room_location_name query string false "room_location_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param player_devices query string false "player_devices"
// @Param control_devices query string false "control_devices"
// @Param schedules query string false "schedules"
// @Param commands query string false "commands"
// @Param sub_type query string false "sub_type"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_item_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-item-info [get]
func Ebcp_exhibition_item_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_item_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_item_info", "id")
}
