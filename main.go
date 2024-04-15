package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	s := NewStorage()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		result, err := s.GetAll(c.Request().Context())
		if err != nil {
			log.Printf("got unexpected error: %v\r\n", err)
			return c.String(http.StatusBadRequest, "unexpected error")
		}
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/:id", func(c echo.Context) error {
		var req GetByIDRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		result, err := s.GetByID(c.Request().Context(), req.ID)
		if err != nil {
			log.Printf("got unexpected error: %v\r\n", err)
			return c.String(http.StatusNotFound, "not found")
		}
		return c.JSON(http.StatusOK, result)
	})

	e.POST("/", func(c echo.Context) error {
		var req TodoCreateRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		result, err := s.Store(c.Request().Context(), req)
		if err != nil {
			log.Printf("got unexpected error: %v\r\n", err)
			return c.String(http.StatusNotFound, "not found")
		}
		return c.JSON(http.StatusOK, result)
	})

	e.DELETE("/:id", func(c echo.Context) error {
		var req GetByIDRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		err := s.Delete(c.Request().Context(), req.ID)
		if err != nil {
			log.Printf("got unexpected error: %v\r\n", err)
			return c.String(http.StatusBadRequest, "unexpected error")
		}
		return c.NoContent(http.StatusOK)
	})

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}

}
