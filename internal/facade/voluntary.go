package facade

import (
	"atados/challenger/internal/constants"
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

func (p *voluntaryFacade) CreateVoluntary(ctx context.Context, voluntaryRequest *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Creating voluntaryRequest", zap.String("correlationID: ", loggerUUID))

	voluntaryVO := voluntaryRequest.ConvertToVoluntaryVO()

	p.logger.Debug("ProductVO", zap.Any("voluntaryRequest", voluntaryVO), zap.String("correlationID: ", loggerUUID))

	voluntaryVO, err := p.voluntaryService.Create(ctx, voluntaryVO)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Voluntary was created", zap.Any("voluntaryRequest", voluntaryVO), zap.String("correlationID: ", loggerUUID))

	response := &dto.CreateVoluntaryResponse{}
	response.ParseFromVoluntaryVO(voluntaryVO)

	p.logger.Debug("Voluntary response", zap.Any("response", response), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Voluntary created", zap.String("correlationID: ", loggerUUID))
	return response, nil
}

func (p *voluntaryFacade) GetVoluntaryByID(ctx context.Context, id uint64) (*dto.GetVoluntaryByIDResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Getting product by ID", zap.String("correlationID: ", loggerUUID))

	p.logger.Debug("Voluntary ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))

	found, voluntary, err := p.voluntaryService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, constants.VoluntaryNotFound
	}

	p.logger.Debug("Voluntary was found", zap.Any("product", voluntary), zap.Bool("found", found), zap.String("correlationID: ", loggerUUID))

	response := &dto.GetVoluntaryByIDResponse{}
	response.ParseFromVoluntaryVO(voluntary)

	p.logger.Debug("Voluntary response", zap.Any("response", response), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Voluntary found", zap.String("correlationID: ", loggerUUID))
	return response, nil
}

func (p *voluntaryFacade) GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllVoluntariesResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Getting all products", zap.String("correlationID: ", loggerUUID))
	products, err := p.voluntaryService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	p.logger.Debug("Voluntaries", zap.Any("products", products), zap.String("correlationID: ", loggerUUID))

	count, err := p.voluntaryService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Count", zap.Int64("count", count), zap.String("correlationID: ", loggerUUID))

	productsResponse := &dto.GetAllVoluntariesResponse{}
	productsResponse.ParseFromVoluntaryVO(products, limit, offset, count)

	p.logger.Debug("Voluntaries response", zap.Any("response", productsResponse), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Voluntaries gotten", zap.String("correlationID: ", loggerUUID))

	return productsResponse, nil
}

func (p *voluntaryFacade) UpdateVoluntary(ctx context.Context, voluntaryID uint64, voluntaryRequest *dto.UpdateVoluntaryRequest) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Updating voluntaryRequest", zap.String("correlationID: ", loggerUUID))
	voluntaryVO := voluntaryRequest.ConvertToVoluntaryVO()
	voluntaryVO.ID = voluntaryID

	p.logger.Debug("ProductVO", zap.Any("voluntaryRequest", voluntaryVO), zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Voluntary ID", zap.Uint64("id", voluntaryID), zap.String("correlationID: ", loggerUUID))

	found, voluntaryFound, err := p.voluntaryService.GetByID(ctx, voluntaryID)
	if err != nil {
		return err
	}
	p.logger.Info("GetVoluntaryByID", zap.String("correlationID: ", loggerUUID))
	if !found {
		return errors.New("voluntaryRequest not found")
	}
	p.logger.Debug("Voluntary found", zap.Any("voluntaryRequest", voluntaryFound), zap.Bool("found", found), zap.String("correlationID: ", loggerUUID))

	voluntaryFound.FirstName = voluntaryVO.FirstName
	voluntaryFound.LastName = voluntaryVO.LastName
	voluntaryFound.Neighborhood = voluntaryVO.Neighborhood
	voluntaryFound.City = voluntaryVO.City

	_, err = p.voluntaryService.Update(ctx, voluntaryFound)
	p.logger.Warn("Voluntary was updated but I'm not using on moment", zap.String("correlationID: ", loggerUUID))
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Voluntary updated", zap.String("correlationID: ", loggerUUID))
	return nil
}

func (p *voluntaryFacade) DeleteVoluntary(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Deleting product", zap.String("correlationID: ", loggerUUID))
	err := p.voluntaryService.Delete(ctx, id)
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Voluntary deleted", zap.String("correlationID: ", loggerUUID))
	return nil
}
