package api

import (
	"ebcp-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitEbcp_holiday_dateRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ebcp-holiday-date/page", Ebcp_holiday_datePageListHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-holiday-date", Ebcp_holiday_dateListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 节假日日期
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param date query string false "date"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param year query string false "year"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ebcp_holiday_date}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-holiday-date/page [get]
func Ebcp_holiday_datePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ebcp_holiday_date](w, r, common.GetDaprClient(), "o_ebcp_holiday_date", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 节假日日期
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param date query string false "date"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param year query string false "year"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ebcp_holiday_date} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ebcp-holiday-date [get]
func Ebcp_holiday_dateListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ebcp_holiday_date](w, r, common.GetDaprClient(), "o_ebcp_holiday_date", "id")
}
