package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_room_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-room-info/page", Ebcp_exhibition_room_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-room-info", Ebcp_exhibition_room_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展厅详细视图，包含展厅信息及其关联的展馆、展览和展项信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param floor query string false "floor"
// @Param floor_value query string false "floor_value"
// @Param floor_name query string false "floor_name"
// @Param location query string false "location"
// @Param location_value query string false "location_value"
// @Param location_name query string false "location_name"
// @Param status query string false "status"
// @Param remarks query string false "remarks"
// @Param hall_id query string false "hall_id"
// @Param hall_name query string false "hall_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param exhibition_status query string false "exhibition_status"
// @Param item_count query string false "item_count"
// @Param items query string false "items"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_exhibition_room_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-room-info/page [get]
func Ebcp_exhibition_room_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_room_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_room_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展厅详细视图，包含展厅信息及其关联的展馆、展览和展项信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param floor query string false "floor"
// @Param floor_value query string false "floor_value"
// @Param floor_name query string false "floor_name"
// @Param location query string false "location"
// @Param location_value query string false "location_value"
// @Param location_name query string false "location_name"
// @Param status query string false "status"
// @Param remarks query string false "remarks"
// @Param hall_id query string false "hall_id"
// @Param hall_name query string false "hall_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param exhibition_status query string false "exhibition_status"
// @Param item_count query string false "item_count"
// @Param items query string false "items"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_room_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-room-info [get]
func Ebcp_exhibition_room_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_room_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_room_info", "id")
}
