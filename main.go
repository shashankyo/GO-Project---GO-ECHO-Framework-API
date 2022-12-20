package main

import (
	"github.com/shashank/golang-echo/cmd/api/service/handlers"
)

func main() {
	e := echo.new()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
