package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthCheckController struct {
}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{}
}

// @Summary healthcheck router
// @Description healthcheck router
// @Tags Healthcheck
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /atados/v1/health [get]
func (healthCheckController) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
