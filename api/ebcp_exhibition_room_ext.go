package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_exhibition_roomExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-room/static-control", StaticControlExhibitionRoomHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-room/start", StartExhibitionRoomHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-room/stop", StopExhibitionRoomHandler)
}

type StaticControlExhibitionRoomRequest struct {
	DeviceIP   string `json:"device_ip"`
	DevicePort int32  `json:"device_port"`
	Command    string `json:"command"`
	Type       string `json:"type"` //start,stop,...
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
}

// @Summary Static control exhibition room
// @Description Static control an exhibition room by ID
// @Tags 展厅
// @Param command_request body StaticControlExhibitionRoomRequest true "Static Control Exhibition Room Request"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-room/static-control [post]
func StaticControlExhibitionRoomHandler(w http.ResponseWriter, r *http.Request) {
	var request StaticControlExhibitionRoomRequest
	err := common.ReadRequestBody(r, &request)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.ControlDeviceCommand(request.DeviceIP, request.DevicePort, request.Command,request.Type,request.DeviceId,request.DeviceType)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK)
}

// @Summary Start exhibition room
// @Description Start an exhibition room by ID
// @Tags 展厅
// @Param room_id query string true "展厅ID"
// @Param type query string true "type,1:数字展项，2:静态展项，不传默认全部"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-room/start [post]
func StartExhibitionRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")
	itemType := r.URL.Query().Get("type")
	go func() {
		err := service.StartExhibitionRoom(roomID, itemType)
		if err != nil {
			common.Logger.Errorf("启动展厅 %s 失败: %v", roomID, err)
		}
	}()
	common.HttpResult(w, common.OK.AppendMsg("后台启动中"))
}

// @Summary Stop exhibition room
// @Description Stop an exhibition room by ID
// @Tags 展厅
// @Param room_id query string true "展厅ID"
// @Param type query string true "type,1:数字展项，2:静态展项，不传默认全部"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-room/stop [post]
func StopExhibitionRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")
	itemType := r.URL.Query().Get("type")
	go func() {
		err := service.StopExhibitionRoom(roomID, itemType)
		if err != nil {
			common.Logger.Errorf("停止展厅 %s 失败: %v", roomID, err)
		}
	}()
	common.HttpResult(w, common.OK.AppendMsg("后台停止中"))
}
