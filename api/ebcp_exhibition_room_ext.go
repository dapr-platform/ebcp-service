package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_exhibition_roomExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-room/static-control", StaticControlExhibitionRoomHandler)
}

type StaticControlExhibitionRoomRequest struct {
	DeviceIP   string `json:"device_ip"`
	DevicePort int32  `json:"device_port"`
	Command    string `json:"command"`
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
	err = service.ControlDeviceCommand(request.DeviceIP, request.DevicePort, request.Command)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
		common.HttpResult(w, common.OK)
}