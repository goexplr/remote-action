package deploy

import (
	"net/http"
	"os"
	"os/exec"
	"remote-action/app/utils"

	"github.com/labstack/echo/v4"
)

// Deploy func
func Deploy(ctx echo.Context) error {
	deployID := ctx.Param("deployId")
	path := os.Getenv(deployID)

	cmd := exec.Command("sh", path)
	res, err := cmd.CombinedOutput()
	if err != nil {
		utils.Logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	utils.Logger.Infof("Perform successfully!")
	return ctx.String(http.StatusOK, string(res))
}
