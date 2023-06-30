package v1

import (
	"atados/challenger/internal/constants"
	"atados/challenger/internal/dto"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type actionController struct {
	actionFacade actionFacade
	logger       *zap.Logger
}

func NewActionController(actionFacade actionFacade, logger *zap.Logger) *actionController {
	return &actionController{
		actionFacade: actionFacade,
		logger:       logger,
	}
}

// @Summary create action router
// @Description create action router
// @Tags Action
// @Accept json
// @Param createActionRequest body dto.CreateActionRequest true "create action"
// @Produce json
// @Success 201 {object} dto.CreateActionResponse
// @Failure 500 {object} error
// @Router /atados/v1/action [post]
func (p *actionController) CreateAction(c echo.Context) error {
	p.logger.Info("Controller: Creating action")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	createAction := &dto.CreateActionRequest{}

	if err := c.Bind(createAction); err != nil {
		p.logger.Error("Error binding request", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("CreateActionRequest", zap.Any("createAction", createAction), zap.String("correlationID: ", loggerUUID))

	action, err := p.actionFacade.CreateAction(ctx, createAction)
	if err != nil {
		p.logger.Error("Error creating action", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Action response", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Action created", zap.String("correlationID: ", loggerUUID))
	return c.JSON(http.StatusCreated, action)
}

// @Summary get action by id router
// @Description get action by id router
// @Tags Action
// @Accept json
// @Param id path int true "id action"
// @Produce json
// @Success 200 {object} dto.GetActionByIDResponse
// @Failure 500 {object} error
// @Router /atados/v1/action/{id} [get]
func (p *actionController) GetActionByID(c echo.Context) error {
	p.logger.Info("Controller: Getting action by ID")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	paramID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("paramID", paramID), zap.String("correlationID: ", loggerUUID))

	action, err := p.actionFacade.GetActionByID(ctx, paramID)
	if err != nil {
		switch {
		case errors.Is(constants.ActionNotFound, err):
			p.logger.Error("Action not found", zap.Error(err), zap.String("correlationID: ", loggerUUID))
			return c.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
		default:
			p.logger.Error("Error getting action", zap.Error(err), zap.String("correlationID: ", loggerUUID))
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	p.logger.Debug("Action response", zap.Any("action", action), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Action found", zap.String("correlationID: ", loggerUUID))
	return c.JSON(http.StatusOK, action)
}

// @Summary get all actions by pagination router
// @Description get all actions by pagination router
// @Tags Action
// @Accept json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Produce json
// @Success 200 {object} dto.GetAllActionsResponse
// @Failure 500 {object} error
// @Router /atados/v1/action [get]
func (p *actionController) GetAllActions(c echo.Context) error {
	p.logger.Info("Controller: Getting all actions")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	p.logger.Debug("Params", zap.Int("limit", limit), zap.Int("offset", offset), zap.String("correlationID: ", loggerUUID))

	actions, err := p.actionFacade.GetAllActions(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Voluntaries response", zap.Any("actions", actions), zap.String("correlationID: ", loggerUUID))

	p.logger.Info("Controller: Getting all actions", zap.String("correlationID: ", loggerUUID))
	return c.JSON(http.StatusOK, actions)
}

// @Summary update action router
// @Description update action router
// @Tags Action
// @Accept json
// @Param id path int true "id action"
// @Param updateActionRequest body dto.UpdateActionRequest true "update action"
// @Produce json
// @Success 200 {string} string "Action updated successfully"
// @Failure 500 {object} error
// @Router /atados/v1/action/{id} [put]
func (p *actionController) UpdateAction(c echo.Context) error {
	p.logger.Info("Controller: Updating action")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	actionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("actionID", actionID), zap.String("correlationID: ", loggerUUID))

	updateAction := &dto.UpdateActionRequest{}
	if err = c.Bind(updateAction); err != nil {
		p.logger.Error("Error binding request", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("UpdateActionRequest", zap.Any("updateAction", updateAction), zap.String("correlationID: ", loggerUUID))
	err = p.actionFacade.UpdateAction(ctx, actionID, updateAction)
	if err != nil {
		p.logger.Error("Error updating action", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Action updated successfully", zap.String("correlationID: ", loggerUUID))
	return c.JSON(http.StatusOK, "Action updated successfully")
}

// @Summary delete action router
// @Description delete action router
// @Tags Action
// @Accept json
// @Param id path int true "id action"
// @Produce json
// @Success 200 {string} string "Action deleted successfully"
// @Failure 500 {object} error
// @Router /atados/v1/action/{id} [delete]
func (p *actionController) DeleteAction(c echo.Context) error {
	p.logger.Info("Controller: Deleting action")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	actionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Param", zap.Uint64("actionID", actionID))
	err = p.actionFacade.DeleteAction(ctx, actionID)
	if err != nil {
		p.logger.Error("Error deleting action", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Action deleted successfully")
	return c.JSON(http.StatusOK, "Action deleted successfully")
}
