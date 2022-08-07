package main

import (
	"os"
	"fmt"

	_ "advisree-be/docs"
	"advisree-be/core"
	"advisree-be/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	defer core.App.Close()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Pre(middleware.Rewrite(map[string]string{
		"/api/*": "/$1",
	}))
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/healthcheck", controllers.HealthCheck)

	port := os.Getenv("API_PORT")
	fmt.Println("port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}