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
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Creating action", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Action", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.Create(ctx, action)
}

func (p *actionService) GetCount(ctx context.Context) (int64, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting count", zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.GetCount(ctx)
}

func (p *actionService) GetByID(ctx context.Context, id uint64) (bool, *model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting action by ID", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.GetByID(ctx, id)
}

func (p *actionService) GetAll(ctx context.Context, limit, offset int) ([]model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting all action", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Limit", zap.Int("limit", limit), zap.Int("offset", offset), zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.GetAll(ctx, limit, offset)
}

func (p *actionService) Update(ctx context.Context, action *model.Action) (*model.Action, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Updating action", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Action", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.Update(ctx, action)
}

func (p *actionService) Delete(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Deleting action", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	return p.actionRepository.Delete(ctx, id)
}
