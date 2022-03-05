package router

import (
	"fmt"
	"os"
	"tests/buisness/controller"

	"github.com/labstack/echo/v4"
)

type router struct {
	server          echo.Echo
	mutanController controller.MutantControllerI
}

func NewRouter(
	server echo.Echo,
	mutanController controller.MutantControllerI,
) *router {
	return &router{
		server,
		mutanController,
	}
}

func (r *router) Start() error {

	r.server.POST("/mutant", r.mutanController.IsMutant)

	return r.server.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
