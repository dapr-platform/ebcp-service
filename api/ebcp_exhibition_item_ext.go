package api

import (
	"ebcp-service/service"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitEbcp_exhibition_itemExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/start", StartExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/pause", PauseExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/{id}/stop", StopExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-start", BatchStartExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-pause", BatchPauseExhibitionItemHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-exhibition-item/batch-stop", BatchStopExhibitionItemHandler)
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


