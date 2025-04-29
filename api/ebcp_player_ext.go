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
	r.Get(common.BASE_CONTEXT+"/ebcp-player/{id}/program-list", GetProgramListHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/fade/{programId}", FadeProgramHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/cut/{programId}", CutProgramHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/pause/{programId}", PauseProgramHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play/{programId}", PlayProgramHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/stop/{programId}", StopProgramHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/play-media/{programId}/{mediaId}", PlayProgaramMediaHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/pause-media/{programId}/{mediaId}", PauseProgramMediaHandler)
	r.Get(common.BASE_CONTEXT+"/ebcp-player/{id}/get-media-process/{programId}/{mediaId}", GetProgramMediaProcessHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/set-media-process/{programId}/{mediaId}", SetProgramMediaProcessHandler)
	// Register sound control routes
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/sound/open", OpenGlobalSoundHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/sound/close", CloseGlobalSoundHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/sound/volume/{volume}", SetGlobalVolumeHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/sound/volume/increase", IncreaseVolumeHandler)
	r.Post(common.BASE_CONTEXT+"/ebcp-player/{id}/sound/volume/decrease", DecreaseVolumeHandler)
}

// @Summary Get program list
// @Description Get program list from specified player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/program-list [get]
func GetProgramListHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	programs, err := service.GetPlayerPrograms(id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK.WithData(programs))
}

// @Summary Fade to program
// @Description Fade to specified program on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/fade/{programId} [post]
func FadeProgramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")

	programId, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err = client.FadeProgram(uint32(programId))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Cut to program
// @Description Cut to specified program on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID, 整数 0,1,2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/cut/{programId} [post]
func CutProgramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")

	programId, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err = client.CutProgram(uint32(programId))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Pause program
// @Description Pause specified program on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/pause/{programId} [post]
func PauseProgramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")

	_, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}

	err = service.PauseProgram(id, programIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Play program
// @Description Play specified program on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID, 整数 0,1,2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play/{programId} [post]
func PlayProgramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")

	_, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}

	err = service.PlayProgram(id, programIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Stop program
// @Description Stop specified program on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID, 整数 0,1,2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/stop/{programId} [post]
func StopProgramHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")

	_, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}

	err = service.StopProgram(id, programIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Play Progaram Media
// @Description Play specified program media on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID,整数id，0，1，2..."
// @Param mediaId path string true "Media ID,整数id，0，1，2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play-media/{programId}/{mediaId} [post]
func PlayProgaramMediaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")
	mediaIdStr := chi.URLParam(r, "mediaId")

	_, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}
	_, err = strconv.ParseUint(mediaIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid media id"))
		return
	}
	err = service.PlayProgramMedia(id, programIdStr, mediaIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Pause Program Media
// @Description Pause specified program media on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID,整数id，0，1，2..."
// @Param mediaId path string true "Media ID,整数id，0，1，2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/pause-media/{programId}/{mediaId} [post]
func PauseProgramMediaHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")
	mediaIdStr := chi.URLParam(r, "mediaId")

	_, err := strconv.ParseUint(programIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid program id"))
		return
	}
	_, err = strconv.ParseUint(mediaIdStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid media id"))
		return
	}
	err = service.PauseProgramMedia(id, programIdStr, mediaIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Get Program Media Process
// @Description Get specified program media process on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID,整数id，0，1，2..."
// @Param mediaId path string true "Media ID,整数id，0，1，2..."
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/get-media-process/{programId}/{mediaId} [get]
func GetProgramMediaProcessHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")
	mediaIdStr := chi.URLParam(r, "mediaId")
	currentTime, totalTime, err := service.GetProgramMediaProcess(id, programIdStr, mediaIdStr)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(map[string]interface{}{"currentTime": currentTime, "totalTime": totalTime}))
}

// @Summary Set Program Media Process
// @Description Set specified program media process on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param programId path string true "Program ID,整数id，0，1，2..."
// @Param mediaId path string true "Media ID,整数id，0，1，2..."
// @Param currentTime query string true "Current Time,整数，单位秒"
// @Param totalTime query string true "Total Time,整数，单位秒"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/set-media-process/{programId}/{mediaId} [post]
func SetProgramMediaProcessHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	programIdStr := chi.URLParam(r, "programId")
	mediaIdStr := chi.URLParam(r, "mediaId")
	currentTimeStr := r.URL.Query().Get("currentTime")
	totalTimeStr := r.URL.Query().Get("totalTime")
	currentTime, err := strconv.ParseUint(currentTimeStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid remain time"))
		return
	}
	totalTime, err := strconv.ParseUint(totalTimeStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid total time"))
		return
	}
	err = service.SetProgramMediaProcess(id, programIdStr, mediaIdStr, int(currentTime), int(totalTime))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Open global sound
// @Description Open global sound on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/sound/open [post]
func OpenGlobalSoundHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err := client.OpenGlobalSound()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Close global sound
// @Description Close global sound on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/sound/close [post]
func CloseGlobalSoundHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err := client.CloseGlobalSound()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Set global volume
// @Description Set global volume on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param volume path string true "Volume level (0-100)"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/sound/volume/{volume} [post]
func SetGlobalVolumeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	volumeStr := chi.URLParam(r, "volume")

	volume, err := strconv.ParseUint(volumeStr, 10, 8)
	if err != nil || volume > 100 {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid volume level (0-100)"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err = client.SetGlobalVolume(uint32(volume))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Increase volume
// @Description Increase volume on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param step query string true "Step (1-100)"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/sound/volume/increase [post]
func IncreaseVolumeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	stepStr := r.URL.Query().Get("step")
	step, err := strconv.ParseUint(stepStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid step"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err = client.IncreaseGlobalVolume(uint32(step))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}

// @Summary Decrease volume
// @Description Decrease volume on player
// @Tags 播放设备
// @Param id path string true "Player ID"
// @Param step query string true "Step (1-100)"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/sound/volume/decrease [post]
func DecreaseVolumeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	stepStr := r.URL.Query().Get("step")
	step, err := strconv.ParseUint(stepStr, 10, 32)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("invalid step"))
		return
	}

	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	err = client.DecreaseGlobalVolume(uint32(step))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}
