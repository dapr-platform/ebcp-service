package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_player_program_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-player-program-info/page", Ebcp_player_program_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-player-program-info", Ebcp_player_program_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 节目详细视图，包含节目信息及其关联的播放设备、展项、展厅和展览信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param program_id query string false "program_id"
// @Param program_index query string false "program_index"
// @Param state query string false "state"
// @Param player_id query string false "player_id"
// @Param player_name query string false "player_name"
// @Param player_ip_address query string false "player_ip_address"
// @Param player_port query string false "player_port"
// @Param player_status query string false "player_status"
// @Param player_current_program_id query string false "player_current_program_id"
// @Param player_current_program_state query string false "player_current_program_state"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param medias query string false "medias"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_player_program_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-program-info/page [get]
func Ebcp_player_program_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_player_program_info](w, r, common.GetDaprClient(), "v_ebcp_player_program_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 节目详细视图，包含节目信息及其关联的播放设备、展项、展厅和展览信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param program_id query string false "program_id"
// @Param program_index query string false "program_index"
// @Param state query string false "state"
// @Param player_id query string false "player_id"
// @Param player_name query string false "player_name"
// @Param player_ip_address query string false "player_ip_address"
// @Param player_port query string false "player_port"
// @Param player_status query string false "player_status"
// @Param player_current_program_id query string false "player_current_program_id"
// @Param player_current_program_state query string false "player_current_program_state"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param medias query string false "medias"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_player_program_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-program-info [get]
func Ebcp_player_program_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_player_program_info](w, r, common.GetDaprClient(), "v_ebcp_player_program_info", "id")
}
