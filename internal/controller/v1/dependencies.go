package v1

import (
	"atados/challenger/internal/dto"
	"context"
)

type voluntaryFacade interface {
	CreateVoluntary(ctx context.Context, voluntary *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryVO, error)
	GetVoluntaryByID(ctx context.Context, id uint64) (*dto.GetVoluntaryByIDResponse, error)
	GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllVoluntariesResponse, error)
	UpdateVoluntary(ctx context.Context, voluntaryID uint64, product *dto.UpdateVoluntaryRequest) error
	DeleteVoluntary(ctx context.Context, id uint64) error
}
