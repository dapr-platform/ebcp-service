package api

import (
	"ebcp-service/service"
	"net/http"
	"strconv"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitEbcp_playerExtRoute(r chi.Router) {
	// Register player control routes
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play-current", PlayCurrentHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play-prev", PlayPrevHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play-next", PlayNextHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play-by-index/{index}", PlayByIndexHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/pause", PauseHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/resume", ResumeHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/set-volume/{volume}", SetVolumeHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/set-window", SetWindowHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/set-visibility/{visible}", SetVisibilityHandler)
}

// @Summary Play current resource
// @Description Play current resource on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play-current [post]
func PlayCurrentHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.PlayCurrent()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Play previous resource
// @Description Play previous resource on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play-prev [post]
func PlayPrevHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.PlayPrev()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Play next resource
// @Description Play next resource on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play-next [post]
func PlayNextHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.PlayNext()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Play resource by index
// @Description Play resource by index on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param index path int true "Resource Index"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play-by-index/{index} [post]
func PlayByIndexHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indexStr := chi.URLParam(r, "index")

	index, err := strconv.ParseUint(indexStr, 10, 64)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid index"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.PlayByIndex(index)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Pause playback
// @Description Pause playback on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/pause [post]
func PauseHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.Pause()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Resume playback
// @Description Resume playback on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/resume [post]
func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, resourceID, err := client.Resume()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success":    success,
		"resourceId": resourceID,
	}))
}

// @Summary Set volume
// @Description Set volume on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param volume path int true "Volume (0-100)"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/set-volume/{volume} [post]
func SetVolumeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	volumeStr := chi.URLParam(r, "volume")

	volume, err := strconv.ParseUint(volumeStr, 10, 64)
	if err != nil || volume > 100 {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid volume"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, err := client.SetVolume(uint16(volume))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success": success,
	}))
}

type WindowParams struct {
	X          int32 `json:"x"`
	Y          int32 `json:"y"`
	Width      int32 `json:"width"`
	Height     int32 `json:"height"`
	Fullscreen bool  `json:"fullscreen"`
}

// @Summary Set window
// @Description Set window position and size on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Accept json
// @Param window body WindowParams true "Window Parameters"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/set-window [post]
func SetWindowHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	var params WindowParams
	if err := common.ReadRequestBody(r, &params); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, err := client.SetWindow(uint16(params.X), uint16(params.Y), uint16(params.Width), uint16(params.Height), params.Fullscreen)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success": success,
	}))
}

// @Summary Set visibility
// @Description Set visibility on specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param visible path bool true "Visibility"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/set-visibility/{visible} [post]
func SetVisibilityHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	visibleStr := chi.URLParam(r, "visible")

	visible, err := strconv.ParseBool(visibleStr)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid visibility value"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	success, err := client.SetVisibility(visible)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{
		"success": success,
	}))
}
