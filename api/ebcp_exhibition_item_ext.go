package api

import (
	"ebcp-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_exhibition_itemExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/start", StartExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/pause", PauseExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/stop", StopExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-start", BatchStartExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-pause", BatchPauseExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-stop", BatchStopExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/static-control", StaticControlExhibitionItemHandler)
}

type StaticControlExhibitionItemRequest struct {
	DeviceIP   string `json:"device_ip"`
	DevicePort int32  `json:"device_port"`
	Command    string `json:"command"`
	Type       string `json:"type"` //start,stop,...
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"` //item,control_device
}

// @Summary Static control exhibition item
// @Description Static control an exhibition item by ID
// @Tags 展项
// @Param command_request body StaticControlExhibitionItemRequest true "Static Control Exhibition Item Request"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/static-control [post]
func StaticControlExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	var request StaticControlExhibitionItemRequest
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

	common.HttpSuccess(w, common.OK)
}

// @Summary Start exhibition item
// @Description Start an exhibition item by ID
// @Tags 展项
// @Param id path string true "Exhibition Item ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/{id}/start [post]
func StartExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := service.StartExhibitionItem(id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Stop exhibition item
// @Description Stop an exhibition item by ID
// @Tags 展项
// @Param id path string true "Exhibition Item ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/{id}/stop [post]
func StopExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := service.StopExhibitionItem(id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Pause exhibition item
// @Description Pause an exhibition item by ID
// @Tags 展项
// @Param id path string true "Exhibition Item ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/{id}/pause [post]
func PauseExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := service.PauseExhibitionItem(id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Batch start exhibition items
// @Description Start multiple exhibition items
// @Tags 展项
// @Accept json
// @Param ids body []string true "Array of Exhibition Item IDs"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/batch-start [post]
func BatchStartExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("no ids provided"))
		return
	}

	err = service.BatchStartExhibitionItems(ids)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Batch stop exhibition items
// @Description Stop multiple exhibition items
// @Tags 展项
// @Accept json
// @Param ids body []string true "Array of Exhibition Item IDs"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/batch-stop [post]
func BatchStopExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("no ids provided"))
		return
	}

	err = service.BatchStopExhibitionItems(ids)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Batch pause exhibition items
// @Description Pause multiple exhibition items
// @Tags 展项
// @Accept json
// @Param ids body []string true "Array of Exhibition Item IDs"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-exhibition-item/batch-pause [post]
func BatchPauseExhibitionItemHandler(w http.ResponseWriter, r *http.Request) {
	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("no ids provided"))
		return
	}

	err = service.BatchPauseExhibitionItems(ids)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}
