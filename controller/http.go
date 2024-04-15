package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/example/config"
	"github.com/vlad-marlo/example/storage"
	"log"
)

type Controller struct {
	echo    *echo.Echo
	storage *storage.Storage
	cfg     *config.Config
}

func New(repo *storage.Storage, cfg *config.Config) *Controller {
	ctrl := &Controller{
		echo:    echo.New(),
		storage: repo,
		cfg:     cfg,
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
	log.Printf("starting http server on address: %s", ctrl.cfg.BindAddr)
	return ctrl.echo.Start(ctrl.cfg.BindAddr)
}
