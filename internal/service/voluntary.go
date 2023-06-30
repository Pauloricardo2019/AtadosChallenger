package service

import (
	"atados/challenger/internal/model"
	"context"
	"go.uber.org/zap"
)

type voluntaryService struct {
	voluntaryRepository voluntaryRepository
	logger              *zap.Logger
}

func NewVoluntaryService(voluntaryRepository voluntaryRepository, logger *zap.Logger) *voluntaryService {
	return &voluntaryService{
		voluntaryRepository: voluntaryRepository,
		logger:              logger,
	}
}

func (p *voluntaryService) Create(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Creating voluntary", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary), zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.Create(ctx, voluntary)
}

func (p *voluntaryService) GetCount(ctx context.Context) (int64, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting count", zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.GetCount(ctx)
}

func (p *voluntaryService) GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting voluntary by ID", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.GetByID(ctx, id)
}

func (p *voluntaryService) GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Getting all voluntaries", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Limit", zap.Int("limit", limit), zap.Int("offset", offset), zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.GetAll(ctx, limit, offset)
}

func (p *voluntaryService) Update(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Updating voluntary", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary), zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.Update(ctx, voluntary)
}

func (p *voluntaryService) Delete(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Service: Deleting voluntary", zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))
	return p.voluntaryRepository.Delete(ctx, id)
}
