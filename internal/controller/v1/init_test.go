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
	ActionControllerMock    *facadeMocks.ActionFacadeMock
}

func setupTestRouter(t *testing.T) (*rest.ServerRest, Facade) {
	t.Helper()

	facades := Facade{
		VoluntaryControllerMock: &facadeMocks.VoluntaryFacadeMock{},
		ActionControllerMock:    &facadeMocks.ActionFacadeMock{},
	}

	cfg := &model.Config{}

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			VoluntaryController:   v1.NewVoluntaryController(facades.VoluntaryControllerMock, logger),
			ActionController:      v1.NewActionController(facades.ActionControllerMock, logger),
			HealthCheckController: v1.NewHealthCheckController(),
		},
	)

	return serverRest, facades

}
