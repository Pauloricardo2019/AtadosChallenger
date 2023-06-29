package facade

import (
	"atados/challenger/internal/model"
	"context"
)

type voluntaryService interface {
	Create(ctx context.Context, product *model.Voluntary) (*model.Voluntary, error)
	GetCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error)
	GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error)
	Update(ctx context.Context, product *model.Voluntary) (*model.Voluntary, error)
	Delete(ctx context.Context, id uint64) error
}
