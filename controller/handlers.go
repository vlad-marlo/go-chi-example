package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/example/model"
	"log"
	"net/http"
)

func (ctrl *Controller) HandleGetAll(c echo.Context) error {
	result, err := ctrl.storage.GetAll(c.Request().Context())
	if err != nil {
		log.Printf("got unexpected error: %v\r\n", err)
		return c.String(http.StatusBadRequest, "unexpected error")
	}
	return c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) HandleGetByID(c echo.Context) error {
	var req model.GetByIDRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := ctrl.storage.GetByID(c.Request().Context(), req.ID)
	if err != nil {
		log.Printf("got unexpected error: %v\r\n", err)
		return c.String(http.StatusNotFound, "not found")
	}
	return c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) HandleCreate(c echo.Context) error {
	var req model.TodoCreateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := ctrl.storage.Store(c.Request().Context(), req)
	if err != nil {
		log.Printf("got unexpected error: %v\r\n", err)
		return c.String(http.StatusNotFound, "not found")
	}
	return c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) HandleDelete(c echo.Context) error {
	var req model.GetByIDRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := ctrl.storage.Delete(c.Request().Context(), req.ID)
	if err != nil {
		log.Printf("got unexpected error: %v\r\n", err)
		return c.String(http.StatusBadRequest, "unexpected error")
	}
	return c.NoContent(http.StatusOK)
}
