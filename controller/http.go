package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/example/storage"
)

type Controller struct {
	echo    *echo.Echo
	storage *storage.Storage
}

func New(repo *storage.Storage) *Controller {
	ctrl := &Controller{
		echo:    echo.New(),
		storage: repo,
	}
	ctrl.configureRoutes()
	return ctrl
}

func (ctrl *Controller) configureRoutes() {
	router := ctrl.echo
	router.GET("/", ctrl.HandleGetAll)
	router.GET("/:id", ctrl.HandleGetByID)
	router.POST("/", ctrl.HandleCreate)
	router.DELETE("/", ctrl.HandleDelete)
}

func (ctrl *Controller) Run() error {
	return ctrl.echo.Start(":8080")
}
