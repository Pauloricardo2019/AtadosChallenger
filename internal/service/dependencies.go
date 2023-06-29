package service

import (
	"atados/challenger/internal/model"
	"context"
)

type voluntaryRepository interface {
	Create(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error)
	GetCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error)
	GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error)
	Update(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error)
	Delete(ctx context.Context, id uint64) error
}
type actionRepository interface {
	Create(ctx context.Context, action *model.Action) (*model.Action, error)
	GetCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.Action, error)
	GetAll(ctx context.Context, limit, offset int) ([]model.Action, error)
	Update(ctx context.Context, action *model.Action) (*model.Action, error)
	Delete(ctx context.Context, id uint64) error
}
