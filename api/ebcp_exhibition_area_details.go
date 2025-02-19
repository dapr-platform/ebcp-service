package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_area_detailsRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area-details/page", Ebcp_exhibition_area_detailsPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-area-details", Ebcp_exhibition_area_detailsListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展区详细视图，包含展区信息及其关联的所有展项信息（JSON格式），展项信息包括名字、状态、类型和备注
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param area_name query string false "area_name"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param current_exhibition_start_time query string false "current_exhibition_start_time"
// @Param current_exhibition_end_time query string false "current_exhibition_end_time"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param location query string false "location"
// @Param exhibition_items query string false "exhibition_items"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_exhibition_area_details}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details/page [get]
func Ebcp_exhibition_area_detailsPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_area_details](w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展区详细视图，包含展区信息及其关联的所有展项信息（JSON格式），展项信息包括名字、状态、类型和备注
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param area_name query string false "area_name"
// @Param current_exhibition_name query string false "current_exhibition_name"
// @Param current_exhibition_start_time query string false "current_exhibition_start_time"
// @Param current_exhibition_end_time query string false "current_exhibition_end_time"
// @Param exhibition_room_id query string false "exhibition_room_id"
// @Param location query string false "location"
// @Param exhibition_items query string false "exhibition_items"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_area_details} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-area-details [get]
func Ebcp_exhibition_area_detailsListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_area_details](w, r, common.GetDaprClient(), "v_ebcp_exhibition_area_details", "id")
}
