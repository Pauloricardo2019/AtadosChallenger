package v1

import (
	"atados/challenger/internal/dto"
	"context"
)

type voluntaryFacade interface {
	CreateVoluntary(ctx context.Context, voluntary *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryResponse, error)
	GetVoluntaryByID(ctx context.Context, id uint64) (*dto.GetVoluntaryByIDResponse, error)
	GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllVoluntariesResponse, error)
	UpdateVoluntary(ctx context.Context, voluntaryID uint64, product *dto.UpdateVoluntaryRequest) error
	DeleteVoluntary(ctx context.Context, id uint64) error
}

type actionFacade interface {
	CreateAction(ctx context.Context, action *dto.CreateActionRequest) (*dto.CreateActionResponse, error)
	GetActionByID(ctx context.Context, id uint64) (*dto.GetActionByIDResponse, error)
	GetAllActions(ctx context.Context, limit, offset int) (*dto.GetAllActionsResponse, error)
	UpdateAction(ctx context.Context, actionID uint64, product *dto.UpdateActionRequest) error
	DeleteAction(ctx context.Context, id uint64) error
}
