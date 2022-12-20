package handlers

import "net/http"

func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK(":1234"))
}
