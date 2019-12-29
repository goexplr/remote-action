package main

import (
	"net/http"
	"remote-action/app/action/deploy"
	"remote-action/app/action/git"

	"github.com/labstack/echo/v4"
)

// Router is manage URL endpoint.
func Router(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Word!")
	})
	e.GET("/git/:action/:repoId", git.Action)
	e.GET("/deploy/:deployId", deploy.Deploy)
}
