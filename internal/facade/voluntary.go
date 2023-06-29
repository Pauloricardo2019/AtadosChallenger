package facade

import (
	"atados/challenger/internal/dto"
	"context"
	"errors"
	"go.uber.org/zap"
)

type voluntaryFacade struct {
	voluntaryService voluntaryService
	logger           *zap.Logger
}

func NewVoluntaryFacade(voluntaryService voluntaryService, logger *zap.Logger) *voluntaryFacade {
	return &voluntaryFacade{
		voluntaryService: voluntaryService,
		logger:           logger,
	}
}

func (p *voluntaryFacade) CreateVoluntary(ctx context.Context, voluntaryRequest *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryVO, error) {
	p.logger.Info("Facade: Creating voluntaryRequest")

	voluntaryVO := voluntaryRequest.ConvertToVoluntaryVO()

	p.logger.Debug("ProductVO", zap.Any("voluntaryRequest", voluntaryVO))

	voluntaryVO, err := p.voluntaryService.Create(ctx, voluntaryVO)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Voluntary was created", zap.Any("voluntaryRequest", voluntaryVO))

	response := &dto.CreateVoluntaryVO{}
	response.ParseFromVoluntaryVO(voluntaryVO)

	p.logger.Debug("Voluntary response", zap.Any("response", response))
	p.logger.Info("Facade: Voluntary created")
	return response, nil
}

func (p *voluntaryFacade) GetVoluntaryByID(ctx context.Context, id uint64) (*dto.GetVoluntaryByIDResponse, error) {
	p.logger.Info("Facade: Getting product by ID")

	p.logger.Debug("Voluntary ID", zap.Uint64("id", id))

	found, productVO, err := p.voluntaryService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("product not found")
	}

	p.logger.Debug("Voluntary was found", zap.Any("product", productVO), zap.Bool("found", found))

	response := &dto.GetVoluntaryByIDResponse{}
	response.ParseFromVoluntaryVO(productVO)

	p.logger.Debug("Voluntary response", zap.Any("response", response))
	p.logger.Info("Facade: Voluntary found")
	return response, nil
}

func (p *voluntaryFacade) GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllVoluntariesResponse, error) {
	p.logger.Info("Facade: Getting all products")
	products, err := p.voluntaryService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	p.logger.Debug("Voluntaries", zap.Any("products", products))

	count, err := p.voluntaryService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Count", zap.Int64("count", count))

	productsResponse := &dto.GetAllVoluntariesResponse{}
	productsResponse.ParseFromProductVO(products, limit, offset, count)

	p.logger.Debug("Voluntaries response", zap.Any("response", productsResponse))
	p.logger.Info("Facade: Voluntaries gotten")

	return productsResponse, nil
}

func (p *voluntaryFacade) UpdateVoluntary(ctx context.Context, voluntaryID uint64, voluntaryRequest *dto.UpdateVoluntaryRequest) error {
	p.logger.Info("Facade: Updating voluntaryRequest")
	voluntaryVO := voluntaryRequest.ConvertToVoluntaryVO()
	voluntaryVO.ID = voluntaryID

	p.logger.Debug("ProductVO", zap.Any("voluntaryRequest", voluntaryVO))
	p.logger.Debug("Voluntary ID", zap.Uint64("id", voluntaryID))

	found, voluntaryFound, err := p.voluntaryService.GetByID(ctx, voluntaryID)
	if err != nil {
		return err
	}
	p.logger.Info("GetVoluntaryByID")
	if !found {
		return errors.New("voluntaryRequest not found")
	}
	p.logger.Debug("Voluntary found", zap.Any("voluntaryRequest", voluntaryFound), zap.Bool("found", found))

	voluntaryFound.FirstName = voluntaryVO.FirstName
	voluntaryFound.LastName = voluntaryVO.LastName
	voluntaryFound.Neighborhood = voluntaryVO.Neighborhood
	voluntaryFound.City = voluntaryVO.City

	_, err = p.voluntaryService.Update(ctx, voluntaryFound)
	p.logger.Warn("Voluntary was updated but I'm not using on moment")
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Voluntary updated")
	return nil
}

func (p *voluntaryFacade) DeleteVoluntary(ctx context.Context, id uint64) error {
	p.logger.Info("Facade: Deleting product")
	err := p.voluntaryService.Delete(ctx, id)
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Voluntary deleted")
	return nil
}
