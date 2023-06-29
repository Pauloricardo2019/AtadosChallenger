package facade

import (
	"atados/challenger/internal/dto"
	"context"
	"go.uber.org/zap"
)

type actionFacade struct {
	actionService actionService
	logger        *zap.Logger
}

func NewActionFacade(actionService actionService, logger *zap.Logger) *actionFacade {
	return &actionFacade{
		actionService: actionService,
		logger:        logger,
	}
}

func (p *actionFacade) CreateAction(ctx context.Context, actionRequest *dto.CreateActionRequest) (*dto.CreateActionVO, error) {
	p.logger.Info("Facade: Creating actionRequest")

	actionVO := actionRequest.ConvertToActionVO()

	p.logger.Debug("ProductVO", zap.Any("actionRequest", actionVO))

	actionVO, err := p.actionService.Create(ctx, actionVO)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Action was created", zap.Any("actionRequest", actionVO))

	response := &dto.CreateActionVO{}
	response.ParseFromActionVO(actionVO)

	p.logger.Debug("Action response", zap.Any("response", response))
	p.logger.Info("Facade: Action created")
	return response, nil
}

func (p *actionFacade) GetActionByID(ctx context.Context, id uint64) (*dto.GetActionByIDResponse, error) {
	p.logger.Info("Facade: Getting product by ID")

	p.logger.Debug("Action ID", zap.Uint64("id", id))

	found, productVO, err := p.actionService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("product not found")
	}

	p.logger.Debug("Action was found", zap.Any("product", productVO), zap.Bool("found", found))

	response := &dto.GetActionByIDResponse{}
	response.ParseFromActionVO(productVO)

	p.logger.Debug("Action response", zap.Any("response", response))
	p.logger.Info("Facade: Action found")
	return response, nil
}

func (p *actionFacade) GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllActionsResponse, error) {
	p.logger.Info("Facade: Getting all products")
	products, err := p.actionService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	p.logger.Debug("Actions", zap.Any("products", products))

	count, err := p.actionService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Count", zap.Int64("count", count))

	productsResponse := &dto.GetAllActionsResponse{}
	productsResponse.ParseFromActionVO(actions, limit, offset, count)

	p.logger.Debug("Actions response", zap.Any("response", productsResponse))
	p.logger.Info("Facade: Actions gotten")

	return productsResponse, nil
}

func (p *actionFacade) UpdateAction(ctx context.Context, actionID uint64, actionRequest *dto.UpdateActionRequest) error {
	p.logger.Info("Facade: Updating actionRequest")
	actionVO := actionRequest.ConvertToActionVO()
	actionVO.ID = actionID

	p.logger.Debug("ProductVO", zap.Any("actionRequest", actionVO))
	p.logger.Debug("Action ID", zap.Uint64("id", actionID))

	found, actionFound, err := p.actionService.GetByID(ctx, actionID)
	if err != nil {
		return err
	}
	p.logger.Info("GetActionByID")
	if !found {
		return errors.New("actionRequest not found")
	}
	p.logger.Debug("Action found", zap.Any("actionRequest", actionFound), zap.Bool("found", found))

	actionFound.FirstName = actionVO.FirstName
	actionFound.LastName = actionVO.LastName
	actionFound.Neighborhood = actionVO.Neighborhood
	actionFound.City = actionVO.City

	_, err = p.actionService.Update(ctx, actionFound)
	p.logger.Warn("Action was updated but I'm not using on moment")
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Action updated")
	return nil
}

func (p *actionFacade) DeleteAction(ctx context.Context, id uint64) error {
	p.logger.Info("Facade: Deleting product")
	err := p.actionService.Delete(ctx, id)
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Action deleted")
	return nil
}
