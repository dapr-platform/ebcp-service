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
	client := service.GetPlayerClient(id)
	if client == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("player not found"))
		return
	}

	programs, err := client.GetProgramList()
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
// @Param programId path string true "Program ID"
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

	err = client.PauseProgram(uint32(programId))
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
// @Param programId path string true "Program ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/play/{programId} [post]
func PlayProgramHandler(w http.ResponseWriter, r *http.Request) {
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

	err = client.PlayProgram(uint32(programId))
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
// @Param programId path string true "Program ID"
// @Produce json
// @Success 200 {object} common.Response "Success"
// @Failure 500 {object} common.Response "Error"
// @Router /ebcp-player/{id}/stop/{programId} [post]
func StopProgramHandler(w http.ResponseWriter, r *http.Request) {
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

	err = client.StopProgram(uint32(programId))
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpSuccess(w, common.OK)
}
