package main

import (
	"net/http"
	"remote-action/app/action/git"

	"github.com/labstack/echo/v4"
)

// Router is manage URL endpoint.
func Router(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Word!")
	})
	g := e.Group("/git")
	g.GET("/:action/:repoId", git.Action)
}
