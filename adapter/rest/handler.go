package rest

import (
	"atados/challenger/docs"
	"atados/challenger/internal/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type Controllers struct {
	HealthCheckController healthCheckController
	VoluntaryController   voluntaryController
	ActionController      actionController
}

type Middlewares struct {
	LogMiddleware logMiddleware
}

type ServerRest struct {
	httpServer  *http.Server
	Engine      *echo.Echo
	config      *model.Config
	controllers Controllers
	middlewares Middlewares
}

func NewRestServer(cfg *model.Config, controllers *Controllers, middlewares *Middlewares) *ServerRest {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	docs.SwaggerInfo.Title = "ATADOS CHALLENGER - API"
	docs.SwaggerInfo.Description = "API CHALLENGER ATADOS"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	server := &ServerRest{
		Engine:      e,
		config:      cfg,
		controllers: *controllers,
		middlewares: *middlewares,
	}

	server.registerRoutes()
	return server
}

func (s *ServerRest) registerRoutes() {
	routeV1 := s.Engine.Group("atados/v1")
	{
		routeV1.GET("/swagger/*", echoSwagger.WrapHandler)

		routeV1.GET("/health", s.controllers.HealthCheckController.HealthCheck)

		voluntaryGroup := routeV1.Group("/voluntary", s.middlewares.LogMiddleware.InitLogger)
		{
			voluntaryGroup.POST("", s.controllers.VoluntaryController.CreateVoluntary)
			voluntaryGroup.GET("/:id", s.controllers.VoluntaryController.GetVoluntaryByID)
			voluntaryGroup.GET("", s.controllers.VoluntaryController.GetAllVoluntaries)
			voluntaryGroup.PUT("/:id", s.controllers.VoluntaryController.UpdateVoluntary)
			voluntaryGroup.DELETE("/:id", s.controllers.VoluntaryController.DeleteVoluntary)
		}

		actionGroup := routeV1.Group("/action", s.middlewares.LogMiddleware.InitLogger)
		{
			actionGroup.POST("", s.controllers.ActionController.CreateAction)
			actionGroup.GET("/:id", s.controllers.ActionController.GetActionByID)
			actionGroup.GET("", s.controllers.ActionController.GetAllActions)
			actionGroup.PUT("/:id", s.controllers.ActionController.UpdateAction)
			actionGroup.DELETE("/:id", s.controllers.ActionController.DeleteAction)
		}
	}
}

func (s *ServerRest) StartListening() {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.RestPort),
		Handler: s.Engine,
	}

	fmt.Println("Listening on port", s.config.RestPort)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err.Error())
	}
}
