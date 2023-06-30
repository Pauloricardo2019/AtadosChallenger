package facade

import (
	"atados/challenger/internal/dto"
	"atados/challenger/internal/mocks"
	"atados/challenger/internal/model"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
	"time"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestVoluntaryFacade_CreateVoluntary(t *testing.T) {
	ctx := context.Background()

	voluntaryServiceMock := &mocks.VoluntaryServiceMock{}

	voluntaryToCreate := &dto.CreateVoluntaryRequest{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryCreated := &model.Voluntary{
		ID:           1,
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	voluntaryServiceMock.On("Create", ctx, mock.Anything).
		Return(
			voluntaryCreated,
			nil,
		)

	voluntaryFacade := NewVoluntaryFacade(voluntaryServiceMock, logger)

	voluntary, err := voluntaryFacade.CreateVoluntary(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID == 1)

}

func TestVoluntaryFacade_GetByIDVoluntary(t *testing.T) {
	ctx := context.Background()

	voluntaryServiceMock := &mocks.VoluntaryServiceMock{}

	voluntaryID := uint64(1)

	voluntaryFound := &model.Voluntary{
		ID:           1,
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	voluntaryServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			voluntaryFound,
			nil,
		)

	voluntaryFacade := NewVoluntaryFacade(voluntaryServiceMock, logger)

	voluntary, err := voluntaryFacade.GetVoluntaryByID(ctx, voluntaryID)
	assert.NoError(t, err)
	assert.True(t, voluntary.ID == 1)
	assert.True(t, voluntary.FirstName == "Miguel")
	assert.True(t, voluntary.LastName == "Ferreira")
	assert.True(t, voluntary.Neighborhood == "Centro")
	assert.True(t, voluntary.City == "Rio de Janeiro")

}

func TestVoluntaryFacade_GetAllVoluntarys(t *testing.T) {
	ctx := context.Background()

	voluntaryServiceMock := &mocks.VoluntaryServiceMock{}

	limit := 10
	offset := 0

	voluntarysFound := []model.Voluntary{
		{
			ID:           1,
			FirstName:    "Miguel",
			LastName:     "Ferreira",
			Neighborhood: "Centro",
			City:         "Rio de Janeiro",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			FirstName:    "Miguel",
			LastName:     "Ferreira",
			Neighborhood: "Centro",
			City:         "Rio de Janeiro",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	voluntaryServiceMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			voluntarysFound,
			nil,
		)

	voluntaryServiceMock.On("GetCount", ctx).
		Return(
			int64(2),
			nil,
		)

	voluntaryFacade := NewVoluntaryFacade(voluntaryServiceMock, logger)

	voluntarys, err := voluntaryFacade.GetAllVoluntaries(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(voluntarys.Voluntaries) == 2)
	assert.True(t, voluntarys.Pagination.Total == 2)

}

func TestVoluntaryFacade_UpdateVoluntary(t *testing.T) {
	ctx := context.Background()

	voluntaryServiceMock := &mocks.VoluntaryServiceMock{}

	voluntaryID := uint64(1)

	voluntaryToUpdated := &dto.UpdateVoluntaryRequest{
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
	}

	voluntaryUpdated := &model.Voluntary{
		ID:           1,
		FirstName:    "Caio",
		LastName:     "Matos",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	voluntaryServiceMock.On("Update", ctx, mock.Anything).
		Return(
			voluntaryUpdated,
			nil,
		)

	voluntaryServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			voluntaryUpdated,
			nil,
		)

	voluntaryFacade := NewVoluntaryFacade(voluntaryServiceMock, logger)

	err := voluntaryFacade.UpdateVoluntary(ctx, voluntaryID, voluntaryToUpdated)
	assert.NoError(t, err)

}

func TestVoluntaryFacade_DeleteVoluntary(t *testing.T) {
	ctx := context.Background()

	voluntaryServiceMock := &mocks.VoluntaryServiceMock{}

	voluntaryID := uint64(1)

	voluntaryServiceMock.On("Delete", ctx, mock.Anything).
		Return(
			nil,
		)

	voluntaryFacade := NewVoluntaryFacade(voluntaryServiceMock, logger)

	err := voluntaryFacade.DeleteVoluntary(ctx, voluntaryID)
	assert.NoError(t, err)

}
