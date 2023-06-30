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

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestActionFacade_CreateAction(t *testing.T) {
	ctx := context.Background()

	actionServiceMock := &mocks.ActionServiceMock{}

	actionToCreate := &dto.CreateActionRequest{
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

	actionServiceMock.On("Create", ctx, mock.Anything).
		Return(
			actionCreated,
			nil,
		)

	actionFacade := NewActionFacade(actionServiceMock, logger)

	action, err := actionFacade.CreateAction(ctx, actionToCreate)
	assert.NoError(t, err)
	assert.True(t, action.ID == 1)

}

func TestActionFacade_GetByIDAction(t *testing.T) {
	ctx := context.Background()

	actionServiceMock := &mocks.ActionServiceMock{}

	actionID := uint64(1)

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

	actionServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			actionFound,
			nil,
		)

	actionFacade := NewActionFacade(actionServiceMock, logger)

	action, err := actionFacade.GetActionByID(ctx, actionID)
	assert.NoError(t, err)
	assert.True(t, action.ID == 1)
	assert.True(t, action.Name == "Action test")
	assert.True(t, action.Institution == "Institution fake")
	assert.True(t, action.City == "São Paulo")
	assert.True(t, action.Neighborhood == "Limoeiro")
	assert.True(t, action.Address == "Rua Palmeira, 25")
	assert.True(t, action.Description == "Reuniao as 15 horas")

}

func TestActionFacade_GetAllActions(t *testing.T) {
	ctx := context.Background()

	actionServiceMock := &mocks.ActionServiceMock{}

	limit := 10
	offset := 0

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

	actionServiceMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			actionsFound,
			nil,
		)

	actionServiceMock.On("GetCount", ctx).
		Return(
			int64(2),
			nil,
		)

	actionFacade := NewActionFacade(actionServiceMock, logger)

	actions, err := actionFacade.GetAllActions(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(actions.Actions) == 2)
	assert.True(t, actions.Pagination.Total == 2)

}

func TestActionFacade_UpdateAction(t *testing.T) {
	ctx := context.Background()

	actionServiceMock := &mocks.ActionServiceMock{}

	actionID := uint64(1)

	actionToUpdated := &dto.UpdateActionRequest{
		Name:         "Action test updated",
		Institution:  "Institution fake updated",
		City:         "Rio de Janeiro",
		Neighborhood: "Bragança",
		Address:      "Rua das flores, 25",
		Description:  "Reuniao as 18 horas",
	}

	actionUpdated := &model.Action{
		ID:           1,
		Name:         "Action test updated",
		Institution:  "Institution fake updated",
		City:         "Rio de Janeiro",
		Neighborhood: "Bragança",
		Address:      "Rua das flores, 25",
		Description:  "Reuniao as 18 horas",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	actionServiceMock.On("Update", ctx, mock.Anything).
		Return(
			actionUpdated,
			nil,
		)

	actionServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			actionUpdated,
			nil,
		)

	actionFacade := NewActionFacade(actionServiceMock, logger)

	err := actionFacade.UpdateAction(ctx, actionID, actionToUpdated)
	assert.NoError(t, err)

}

func TestActionFacade_DeleteAction(t *testing.T) {
	ctx := context.Background()

	actionServiceMock := &mocks.ActionServiceMock{}

	actionID := uint64(1)

	actionServiceMock.On("Delete", ctx, mock.Anything).
		Return(
			nil,
		)

	actionFacade := NewActionFacade(actionServiceMock, logger)

	err := actionFacade.DeleteAction(ctx, actionID)
	assert.NoError(t, err)

}
