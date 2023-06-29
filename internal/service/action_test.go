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

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestActionService_Create(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	actionToCreate := &model.Action{
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
	}

	actionCreated := &model.Action{
		ID:           1,
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	actionRepositoryMock.On("Create", ctx, mock.Anything).
		Return(
			actionCreated,
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	actionCreated, err := actionService.Create(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, actionCreated.ID == 1)

}

func TestActionService_GetCount(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	actionRepositoryMock.On("GetCount", ctx).
		Return(
			int64(1),
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	count, err := actionService.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count == 1)

}

func TestActionService_GetByID(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	idMock := uint64(1)

	actionFound := &model.Action{
		ID:           1,
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	actionRepositoryMock.On("GetByID", ctx, idMock).
		Return(
			true,
			actionFound,
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	found, actionFound, err := actionService.GetByID(ctx, idMock)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, actionFound.ID == 1)

}

func TestActionService_GetAll(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	actionsFound := []model.Action{
		{
			ID:           1,
			Name:         "Action test",
			Institution:  "Institution fake",
			City:         "São Paulo",
			Neighborhood: "Limoeiro",
			Address:      "Rua Palmeira, 25",
			Description:  "Reuniao as 15 horas",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "Action test",
			Institution:  "Institution fake",
			City:         "São Paulo",
			Neighborhood: "Limoeiro",
			Address:      "Rua Palmeira, 25",
			Description:  "Reuniao as 15 horas",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	actionRepositoryMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			actionsFound,
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	actions, err := actionService.GetAll(ctx, 1, 10)
	assert.NoError(t, err)
	assert.True(t, len(actions) == 2)

}

func TestActionService_Update(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	actionToUpdate := &model.Action{
		ID:           1,
		Name:         "Action test",
		Institution:  "Institution fake",
		City:         "São Paulo",
		Neighborhood: "Limoeiro",
		Address:      "Rua Palmeira, 25",
		Description:  "Reuniao as 15 horas",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	actionUpdated := &model.Action{
		ID:           1,
		Name:         "Action test updated",
		Institution:  "Institution fake updated",
		City:         "Rio de Janeiro",
		Neighborhood: "Mangueira",
		Address:      "Rua das flores, 25",
		Description:  "Reuniao as 19 horas",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	actionRepositoryMock.On("Update", ctx, mock.Anything).
		Return(
			actionUpdated,
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	actionUpdated, err := actionService.Update(ctx, actionToUpdate)
	assert.NoError(t, err)
	assert.True(t, actionUpdated.ID == 1)
	assert.True(t, actionUpdated.Name == "Action test updated")
	assert.True(t, actionUpdated.Institution == "Institution fake updated")
	assert.True(t, actionUpdated.City == "Rio de Janeiro")
	assert.True(t, actionUpdated.Neighborhood == "Mangueira")
	assert.True(t, actionUpdated.Address == "Rua das flores, 25")
	assert.True(t, actionUpdated.Description == "Reuniao as 19 horas")

}

func TestActionService_Delete(t *testing.T) {
	ctx := context.Background()

	actionRepositoryMock := &mocks.ActionRepositoryMock{}

	actionID := uint64(1)

	actionRepositoryMock.On("Delete", ctx, actionID).
		Return(
			nil,
		)

	actionService := NewActionService(actionRepositoryMock, logger)

	err := actionService.Delete(ctx, actionID)
	assert.NoError(t, err)
}
