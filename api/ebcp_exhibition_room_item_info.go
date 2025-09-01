package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_room_item_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-room-item-info/page", Ebcp_exhibition_room_item_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-room-item-info", Ebcp_exhibition_room_item_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展览展厅展项信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param status query string false "status"
// @Param total_room_count query string false "total_room_count"
// @Param total_item_count query string false "total_item_count"
// @Param rooms query string false "rooms"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_exhibition_room_item_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-room-item-info/page [get]
func Ebcp_exhibition_room_item_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_room_item_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_room_item_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展览展厅展项信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param status query string false "status"
// @Param total_room_count query string false "total_room_count"
// @Param total_item_count query string false "total_item_count"
// @Param rooms query string false "rooms"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_room_item_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-room-item-info [get]
func Ebcp_exhibition_room_item_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_room_item_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_room_item_info", "id")
}
