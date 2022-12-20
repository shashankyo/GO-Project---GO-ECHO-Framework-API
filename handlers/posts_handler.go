package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/shashank/golang-echo/cmd/api/service"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
