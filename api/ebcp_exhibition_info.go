package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-info/page", Ebcp_exhibition_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-info", Ebcp_exhibition_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展览信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param hall_id query string false "hall_id"
// @Param hall_name query string false "hall_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_location query string false "room_location"
// @Param room_status query string false "room_status"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param item_type query string false "item_type"
// @Param item_status query string false "item_status"
// @Param room_count query string false "room_count"
// @Param item_count query string false "item_count"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_exhibition_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-info/page [get]
func Ebcp_exhibition_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_info", "exhibition_id")

}

// @Summary query objects
// @Description query objects
// @Tags 展览信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param exhibition_id query string false "exhibition_id"
// @Param exhibition_name query string false "exhibition_name"
// @Param exhibition_start_time query string false "exhibition_start_time"
// @Param exhibition_end_time query string false "exhibition_end_time"
// @Param hall_id query string false "hall_id"
// @Param hall_name query string false "hall_name"
// @Param room_id query string false "room_id"
// @Param room_name query string false "room_name"
// @Param room_floor query string false "room_floor"
// @Param room_location query string false "room_location"
// @Param room_status query string false "room_status"
// @Param item_id query string false "item_id"
// @Param item_name query string false "item_name"
// @Param item_type query string false "item_type"
// @Param item_status query string false "item_status"
// @Param room_count query string false "room_count"
// @Param item_count query string false "item_count"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-info [get]
func Ebcp_exhibition_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_info", "exhibition_id")
}
