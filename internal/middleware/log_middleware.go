package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type logMiddleware struct{}

func NewLogMiddleware() *logMiddleware {
	return &logMiddleware{}
}

func (l *logMiddleware) InitLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logID := uuid.New()
		c.Set("logger", logID.String())
		return next(c)
	}
}
