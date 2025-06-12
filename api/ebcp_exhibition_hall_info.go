package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_exhibition_hall_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-hall-info/page", Ebcp_exhibition_hall_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-exhibition-hall-info", Ebcp_exhibition_hall_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 展馆详细视图，包含展馆信息及其关联的展厅和展项信息（JSON格式）
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param _select query string true "_select"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param remarks query string false "remarks"
// @Param rooms query string false "rooms"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Ebcp_exhibition_hall_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-hall-info/page [get]
func Ebcp_exhibition_hall_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_exhibition_hall_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_hall_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 展馆详细视图，包含展馆信息及其关联的展厅和展项信息（JSON格式）
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param remarks query string false "remarks"
// @Param rooms query string false "rooms"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_exhibition_hall_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition-hall-info [get]
func Ebcp_exhibition_hall_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_exhibition_hall_info](w, r, common.GetDaprClient(), "v_ebcp_exhibition_hall_info", "id")
}
