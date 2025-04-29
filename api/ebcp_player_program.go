package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_player_programRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-player-program/page", Ebcp_player_programPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-player-program", Ebcp_player_programListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 播放设备节目
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param player_id query string false "player_id"
// @Param program_id query string false "program_id"
// @Param program_index query string false "program_index"
// @Param state query string false "state"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_player_program}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-program/page [get]
func Ebcp_player_programPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_player_program](w, r, common.GetDaprClient(), "o_ebcp_player_program", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 播放设备节目
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param player_id query string false "player_id"
// @Param program_id query string false "program_id"
// @Param program_index query string false "program_index"
// @Param state query string false "state"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_player_program} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-player-program [get]
func Ebcp_player_programListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_player_program](w, r, common.GetDaprClient(), "o_ebcp_player_program", "id")
}
