package facade

import (
	"atados/challenger/internal/constants"
	"atados/challenger/internal/dto"
	"context"
	"errors"
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

func (p *actionFacade) CreateAction(ctx context.Context, actionRequest *dto.CreateActionRequest) (*dto.CreateActionResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Creating actionRequest", zap.String("correlationID: ", loggerUUID))

	actionVO := actionRequest.ConvertToActionVO()

	p.logger.Debug("ActionVO", zap.Any("actionRequest", actionVO), zap.String("correlationID: ", loggerUUID))

	actionVO, err := p.actionService.Create(ctx, actionVO)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Action was created", zap.Any("actionRequest", actionVO), zap.String("correlationID: ", loggerUUID))

	response := &dto.CreateActionResponse{}
	response.ParseFromActionVO(actionVO)

	p.logger.Debug("Action response", zap.Any("response", response), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Action created", zap.String("correlationID: ", loggerUUID))
	return response, nil
}

func (p *actionFacade) GetActionByID(ctx context.Context, id uint64) (*dto.GetActionByIDResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Getting action by ID", zap.String("correlationID: ", loggerUUID))

	p.logger.Debug("Action ID", zap.Uint64("id", id), zap.String("correlationID: ", loggerUUID))

	found, actionVO, err := p.actionService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, constants.ActionNotFound
	}

	p.logger.Debug("Action was found", zap.Any("action", actionVO), zap.Bool("found", found), zap.String("correlationID: ", loggerUUID))

	response := &dto.GetActionByIDResponse{}
	response.ParseFromActionVO(actionVO)

	p.logger.Debug("Action response", zap.Any("response", response), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Action found", zap.String("correlationID: ", loggerUUID))
	return response, nil
}

func (p *actionFacade) GetAllActions(ctx context.Context, limit, offset int) (*dto.GetAllActionsResponse, error) {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Getting all actions", zap.String("correlationID: ", loggerUUID))
	actions, err := p.actionService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	p.logger.Debug("Actions", zap.Any("actions", actions), zap.String("correlationID: ", loggerUUID))

	count, err := p.actionService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Count", zap.Int64("count", count), zap.String("correlationID: ", loggerUUID))

	actionsResponse := &dto.GetAllActionsResponse{}
	actionsResponse.ParseFromActionVO(actions, limit, offset, count)

	p.logger.Debug("Actions response", zap.Any("response", actionsResponse), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Facade: Actions gotten", zap.String("correlationID: ", loggerUUID))

	return actionsResponse, nil
}

func (p *actionFacade) UpdateAction(ctx context.Context, actionID uint64, actionRequest *dto.UpdateActionRequest) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Updating actionRequest", zap.String("correlationID: ", loggerUUID))
	actionVO := actionRequest.ConvertToActionVO()
	actionVO.ID = actionID

	p.logger.Debug("ActionVO", zap.Any("actionRequest", actionVO), zap.String("correlationID: ", loggerUUID))
	p.logger.Debug("Action ID", zap.Uint64("id", actionID), zap.String("correlationID: ", loggerUUID))

	found, actionFound, err := p.actionService.GetByID(ctx, actionID)
	if err != nil {
		return err
	}
	p.logger.Info("GetActionByID", zap.String("correlationID: ", loggerUUID))
	if !found {
		return errors.New("actionRequest not found")
	}
	p.logger.Debug("Action found", zap.Any("actionRequest", actionFound), zap.Bool("found", found), zap.String("correlationID: ", loggerUUID))

	actionFound.Name = actionVO.Name
	actionFound.Institution = actionVO.Institution
	actionFound.City = actionVO.City
	actionFound.Neighborhood = actionVO.Neighborhood
	actionFound.Address = actionVO.Address
	actionFound.Description = actionVO.Description

	_, err = p.actionService.Update(ctx, actionFound)
	p.logger.Warn("Action was updated but I'm not using on moment", zap.String("correlationID: ", loggerUUID))
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Action updated", zap.String("correlationID: ", loggerUUID))
	return nil
}

func (p *actionFacade) DeleteAction(ctx context.Context, id uint64) error {
	loggerUUID := ctx.Value("logger").(string)
	p.logger.Info("Facade: Deleting action", zap.String("correlationID: ", loggerUUID))
	err := p.actionService.Delete(ctx, id)
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Action deleted", zap.String("correlationID: ", loggerUUID))
	return nil
}
