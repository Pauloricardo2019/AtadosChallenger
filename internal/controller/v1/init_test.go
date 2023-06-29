package v1_test

import (
	"atados/challenger/adapter/rest"
	"atados/challenger/internal/controller/v1"
	facadeMocks "atados/challenger/internal/mocks"
	"atados/challenger/internal/model"
	"go.uber.org/zap"
	"testing"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

type Facade struct {
	VoluntaryControllerMock *facadeMocks.VoluntaryFacadeMock
}

func setupTestRouter(t *testing.T) (*rest.ServerRest, Facade) {
	t.Helper()

	facades := Facade{
		VoluntaryControllerMock: &facadeMocks.VoluntaryFacadeMock{},
	}

	cfg := &model.Config{}

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			VoluntaryController:   v1.NewVoluntaryController(facades.VoluntaryControllerMock, logger),
			HealthCheckController: v1.NewHealthCheckController(),
		},
	)

	return serverRest, facades

}
