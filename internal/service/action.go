package service

import (
	"atados/challenger/internal/model"
	"context"
	"go.uber.org/zap"
)

type actionService struct {
	actionRepository actionRepository
	logger           *zap.Logger
}

func NewActionService(actionRepository actionRepository, logger *zap.Logger) *actionService {
	return &actionService{
		actionRepository: actionRepository,
		logger:           logger,
	}
}

func (p *actionService) Create(ctx context.Context, action *model.Action) (*model.Action, error) {
	p.logger.Info("Service: Creating action")
	p.logger.Debug("Action", zap.Any("action", action))
	return p.actionRepository.Create(ctx, action)
}

func (p *actionService) GetCount(ctx context.Context) (int64, error) {
	p.logger.Info("Service: Getting count")
	return p.actionRepository.GetCount(ctx)
}

func (p *actionService) GetByID(ctx context.Context, id uint64) (bool, *model.Action, error) {
	p.logger.Info("Service: Getting action by ID")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.actionRepository.GetByID(ctx, id)
}

func (p *actionService) GetAll(ctx context.Context, limit, offset int) ([]model.Action, error) {
	p.logger.Info("Service: Getting all action")
	p.logger.Debug("Limit", zap.Int("limit", limit), zap.Int("offset", offset))
	return p.actionRepository.GetAll(ctx, limit, offset)
}

func (p *actionService) Update(ctx context.Context, action *model.Action) (*model.Action, error) {
	p.logger.Info("Service: Updating action")
	p.logger.Debug("Action", zap.Any("action", action))
	return p.actionRepository.Update(ctx, action)
}

func (p *actionService) Delete(ctx context.Context, id uint64) error {
	p.logger.Info("Service: Deleting action")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.actionRepository.Delete(ctx, id)
}
