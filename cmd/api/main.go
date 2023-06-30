package main

import (
	"atados/challenger/adapter/provider"
	"atados/challenger/adapter/repository"
	"atados/challenger/adapter/rest"
	"atados/challenger/internal/config"
	"atados/challenger/internal/controller/v1"
	"atados/challenger/internal/facade"
	"atados/challenger/internal/middleware"
	"atados/challenger/internal/service"
	"go.uber.org/zap"
	"time"
)

var envPath = "./dev.env"

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	logger := zap.NewExample()

	defer logger.Sync()

	//Config
	cfg := config.NewConfig(logger).GetConfig(envPath)

	logger.Info("Setup config",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Database Connection
	dbProvider := provider.NewDatabaseProvider(cfg, logger)
	dbConn, err := dbProvider.GetConnection()
	if err != nil {
		panic(err)
		return
	}

	logger.Info("Setup database",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Repositories
	voluntaryRepository := repository.NewVoluntaryRepository(dbConn, logger)
	actionRepository := repository.NewActionRepository(dbConn, logger)

	logger.Info("Setup repositories",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Services
	voluntaryService := service.NewVoluntaryService(voluntaryRepository, logger)
	actionService := service.NewActionService(actionRepository, logger)

	logger.Info("Setup services",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Facades
	voluntaryFacade := facade.NewVoluntaryFacade(voluntaryService, logger)
	actionFacade := facade.NewActionFacade(actionService, logger)

	logger.Info("Setup facades",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Middleware
	logMiddleware := middleware.NewLogMiddleware()

	//Get Controllers
	voluntaryController := v1.NewVoluntaryController(voluntaryFacade, logger)
	actionController := v1.NewActionController(actionFacade, logger)
	healthCheckController := v1.NewHealthCheckController()

	logger.Info("Setup controllers",
		zap.Time("StartedAt", time.Now()),
	)

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			HealthCheckController: healthCheckController,
			VoluntaryController:   voluntaryController,
			ActionController:      actionController,
		},
		&rest.Middlewares{
			LogMiddleware: logMiddleware,
		},
	)

	logger.Info("Setup server",
		zap.Time("StartedAt", time.Now()),
	)

	serverRest.StartListening()

}
