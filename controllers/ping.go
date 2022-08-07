package controllers

import (
	"advisree-be/core"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HealthCheck check API health status
// @Summary Check API health status
// @Description Check API health status
// @Tags healthcheck
// @ID healthcheck-healthcheck
// @Accept json
// @Produce application/json
// @Success 200 {object} interface{}
// @Router /healthcheck [get]
func HealthCheck(c echo.Context) error {
	defer c.Request().Body.Close()

	response := map[string]interface{}{
		"time": fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05")),
		"db":   true,
	}

	if err := core.App.DB.DB().Ping(); err != nil {
		response["db"] = false
	}

	return c.JSON(http.StatusOK, response)
}