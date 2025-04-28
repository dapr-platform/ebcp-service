package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_player_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-player-info/page", Ebcp_player_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-player-info", Ebcp_player_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 播放设备详细视图，包含设备信息及其关联的展项、展厅、展览和节目信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param status query string false "status"
// @Param current_program_id query string false "current_program_id"
// @Param current_program_state query string false "current_program_state"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param item_type query string false "item_type"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param programs query string false "programs"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_player_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-info/page [get]
func Ebcp_player_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_player_info](w, r, common.GetDaprClient(), "v_ebcp_player_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 播放设备详细视图，包含设备信息及其关联的展项、展厅、展览和节目信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param ip_address query string false "ip_address"
// @Param port query string false "port"
// @Param version query string false "version"
// @Param status query string false "status"
// @Param current_program_id query string false "current_program_id"
// @Param current_program_state query string false "current_program_state"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param item_type query string false "item_type"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_floor_value query string false "room_floor_value"
// @Param room_floor_name query string false "room_floor_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param programs query string false "programs"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_player_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-info [get]
func Ebcp_player_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_player_info](w, r, common.GetDaprClient(), "v_ebcp_player_info", "id")
}
