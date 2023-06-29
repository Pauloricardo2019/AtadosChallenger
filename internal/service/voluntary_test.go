package service

import (
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

func TestVoluntaryService_Create(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

	voluntaryToCreate := &model.Voluntary{
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

	voluntaryRepositoryMock.On("Create", ctx, mock.Anything).
		Return(
			voluntaryCreated,
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	voluntaryCreated, err := voluntaryService.Create(ctx, voluntaryToCreate)
	assert.NoError(t, err)
	assert.True(t, voluntaryCreated.ID == 1)

}

func TestVoluntaryService_GetCount(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

	voluntaryRepositoryMock.On("GetCount", ctx).
		Return(
			int64(1),
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	count, err := voluntaryService.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count == 1)

}

func TestVoluntaryService_GetByID(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

	idMock := uint64(1)

	voluntaryFound := &model.Voluntary{
		ID:           1,
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	voluntaryRepositoryMock.On("GetByID", ctx, idMock).
		Return(
			true,
			voluntaryFound,
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	found, voluntaryFound, err := voluntaryService.GetByID(ctx, idMock)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, voluntaryFound.ID == 1)

}

func TestVoluntaryService_GetAll(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

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
			FirstName:    "Caio",
			LastName:     "Matos",
			Neighborhood: "Centro",
			City:         "Rio de Janeiro",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	voluntaryRepositoryMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			voluntarysFound,
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	voluntarys, err := voluntaryService.GetAll(ctx, 1, 10)
	assert.NoError(t, err)
	assert.True(t, len(voluntarys) == 2)

}

func TestVoluntaryService_Update(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

	voluntaryToUpdate := &model.Voluntary{
		ID:           1,
		FirstName:    "Miguel",
		LastName:     "Ferreira",
		Neighborhood: "Centro",
		City:         "Rio de Janeiro",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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

	voluntaryRepositoryMock.On("Update", ctx, mock.Anything).
		Return(
			voluntaryUpdated,
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	voluntaryUpdated, err := voluntaryService.Update(ctx, voluntaryToUpdate)
	assert.NoError(t, err)
	assert.True(t, voluntaryUpdated.ID == 1)
	assert.True(t, voluntaryUpdated.FirstName == "Caio")
	assert.True(t, voluntaryUpdated.LastName == "Matos")
}

func TestVoluntaryService_Delete(t *testing.T) {
	ctx := context.Background()

	voluntaryRepositoryMock := &mocks.VoluntaryRepositoryMock{}

	voluntaryID := uint64(1)

	voluntaryRepositoryMock.On("Delete", ctx, voluntaryID).
		Return(
			nil,
		)

	voluntaryService := NewVoluntaryService(voluntaryRepositoryMock, logger)

	err := voluntaryService.Delete(ctx, voluntaryID)
	assert.NoError(t, err)
}
