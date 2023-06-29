package rest

import "github.com/labstack/echo/v4"

type (
	healthCheckController interface {
		HealthCheck(c echo.Context) error
	}

	voluntaryController interface {
		CreateVoluntary(c echo.Context) error
		GetVoluntaryByID(c echo.Context) error
		GetAllVoluntaries(c echo.Context) error
		UpdateVoluntary(c echo.Context) error
		DeleteVoluntary(c echo.Context) error
	}
)
