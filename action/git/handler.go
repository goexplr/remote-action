package git

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/exec"
	"remote-action/app/utils"
)

// ACTIONS supported
var ACTIONS = []string{"fetch", "pull", "log", "status"}

// Action func
func Action(ctx echo.Context) error {
	action := ctx.Param("action")
	found := false
	for _, act := range ACTIONS {
		if act == action {
			found = true
		}
	}
	if !found {
		return ctx.String(http.StatusBadRequest, "Action not supported")
	}
	repoID := ctx.Param("repoId")
	path := os.Getenv(repoID)

	cmd := exec.Command("git", action)
	cmd.Dir = path

	res, err := cmd.Output()
	if err != nil {
		utils.Logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	utils.Logger.Infof("Perform successfully!")
	return ctx.String(http.StatusOK, string(res))
}
