package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_exhibition_extRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition/start", Ebcp_exhibition_extStartHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition/stop", Ebcp_exhibition_extStopHandler)
}

// @Summary 展览一键启动
// @Description 展览一键启动
// @Tags 展览
// @Accept  json
// @Produce  json
// @Param exhibition_id query string true "展览ID"
// @Param type query string true "type,1:数字展项，2:静态展项，不传默认全部"
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition/start [post]
func Ebcp_exhibition_extStartHandler(w http.ResponseWriter, r *http.Request) {
	itemType := r.URL.Query().Get("type")
	exhibitionId := r.URL.Query().Get("exhibition_id")
	go func() {
		err := service.StartExhibition(exhibitionId, itemType)
		if err != nil {
			common.Logger.Errorf("启动展览 %s 失败: %v", exhibitionId, err)
		}
	}()

	common.HttpResult(w, common.OK.AppendMsg("后台启动中"))
}

// @Summary 展览一键停止
// @Description 展览一键停止
// @Tags 展览
// @Accept  json
// @Produce  json
// @Param exhibition_id query string true "展览ID"
// @Param type query string true "type,1:数字展项，2:静态展项，不传默认全部"
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ebcp-exhibition/stop [post]
func Ebcp_exhibition_extStopHandler(w http.ResponseWriter, r *http.Request) {
	itemType := r.URL.Query().Get("type")
	exhibitionId := r.URL.Query().Get("exhibition_id")
	go func() {
		err := service.StopExhibition(exhibitionId, itemType)
		if err != nil {
			common.Logger.Errorf("停止展览 %s 失败: %v", exhibitionId, err)
		}
	}()
	common.HttpResult(w, common.OK.AppendMsg("后台停止中"))
}
