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
	p.logger.Info("Service: Creating voluntary")
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary))
	return p.voluntaryRepository.Create(ctx, voluntary)
}

func (p *voluntaryService) GetCount(ctx context.Context) (int64, error) {
	p.logger.Info("Service: Getting count")
	return p.voluntaryRepository.GetCount(ctx)
}

func (p *voluntaryService) GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error) {
	p.logger.Info("Service: Getting voluntary by ID")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.voluntaryRepository.GetByID(ctx, id)
}

func (p *voluntaryService) GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error) {
	p.logger.Info("Service: Getting all voluntaries")
	p.logger.Debug("Limit", zap.Int("limit", limit), zap.Int("offset", offset))
	return p.voluntaryRepository.GetAll(ctx, limit, offset)
}

func (p *voluntaryService) Update(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	p.logger.Info("Service: Updating voluntary")
	p.logger.Debug("Voluntary", zap.Any("voluntary", voluntary))
	return p.voluntaryRepository.Update(ctx, voluntary)
}

func (p *voluntaryService) Delete(ctx context.Context, id uint64) error {
	p.logger.Info("Service: Deleting voluntary")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.voluntaryRepository.Delete(ctx, id)
}
