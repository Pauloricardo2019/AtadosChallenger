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

type voluntaryController struct {
	voluntaryFacade voluntaryFacade
	logger          *zap.Logger
}

func NewVoluntaryController(voluntaryFacade voluntaryFacade, logger *zap.Logger) *voluntaryController {
	return &voluntaryController{
		voluntaryFacade: voluntaryFacade,
		logger:          logger,
	}
}

// @Summary create voluntary router
// @Description create voluntary router
// @Tags Voluntary
// @Accept json
// @Param createVoluntaryRequest body dto.CreateVoluntaryRequest true "create voluntary"
// @Produce json
// @Success 201 {object} dto.CreateVoluntaryResponse
// @Failure 500 {object} error
// @Router /atados/v1/voluntary [post]
func (p *voluntaryController) CreateVoluntary(c echo.Context) error {
	p.logger.Info("Controller: Creating voluntary")

	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	createVoluntary := &dto.CreateVoluntaryRequest{}

	if err := c.Bind(createVoluntary); err != nil {
		p.logger.Error("Error binding request", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("CreateVoluntaryRequest", zap.Any("createVoluntary", createVoluntary), zap.String("correlationID: ", loggerUUID))

	voluntary, err := p.voluntaryFacade.CreateVoluntary(ctx, createVoluntary)
	if err != nil {
		p.logger.Error("Error creating voluntary", zap.Error(err), zap.String("correlationID: ", loggerUUID))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Voluntary response", zap.Any("voluntary", voluntary), zap.String("correlationID: ", loggerUUID))
	p.logger.Info("Voluntary created", zap.String("correlationID: ", loggerUUID))
	return c.JSON(http.StatusCreated, voluntary)
}

// @Summary get voluntary by id router
// @Description get voluntary by id router
// @Tags Voluntary
// @Accept json
// @Param id path int true "id voluntary"
// @Produce json
// @Success 200 {object} dto.GetVoluntaryByIDResponse
// @Failure 500 {object} error
// @Router /atados/v1/voluntary/{id} [get]
func (p *voluntaryController) GetVoluntaryByID(c echo.Context) error {
	p.logger.Info("Controller: Getting voluntary by ID")

	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	paramID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("paramID", paramID))

	voluntary, err := p.voluntaryFacade.GetVoluntaryByID(ctx, paramID)
	if err != nil {
		switch {
		case errors.Is(constants.VoluntaryNotFound, err):
			p.logger.Error("voluntary not found", zap.Error(err))
			return c.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
		default:
			p.logger.Error("Error getting voluntary", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	p.logger.Debug("Voluntary response", zap.Any("voluntary", voluntary))
	p.logger.Info("Voluntary found")
	return c.JSON(http.StatusOK, voluntary)
}

// @Summary get all voluntarys by pagination router
// @Description get all voluntarys by pagination router
// @Tags Voluntary
// @Accept json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Produce json
// @Success 200 {object} dto.GetAllVoluntariesResponse
// @Failure 500 {object} error
// @Router /atados/v1/voluntary [get]
func (p *voluntaryController) GetAllVoluntaries(c echo.Context) error {
	p.logger.Info("Controller: Getting all voluntaries")
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

	p.logger.Debug("Params", zap.Int("limit", limit), zap.Int("offset", offset))

	voluntaries, err := p.voluntaryFacade.GetAllVoluntaries(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Voluntaries response", zap.Any("voluntaries", voluntaries))

	p.logger.Info("Controller: Getting all voluntaries")
	return c.JSON(http.StatusOK, voluntaries)
}

// @Summary update voluntary router
// @Description update voluntary router
// @Tags Voluntary
// @Accept json
// @Param id path int true "id voluntary"
// @Param updateVoluntaryRequest body dto.UpdateVoluntaryRequest true "update voluntary"
// @Produce json
// @Success 200 {string} string "Voluntary updated successfully"
// @Failure 500 {object} error
// @Router /atados/v1/voluntary/{id} [put]
func (p *voluntaryController) UpdateVoluntary(c echo.Context) error {
	p.logger.Info("Controller: Updating voluntary")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	voluntaryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("voluntaryID", voluntaryID))

	updateVoluntary := &dto.UpdateVoluntaryRequest{}
	if err = c.Bind(updateVoluntary); err != nil {
		p.logger.Error("Error binding request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("UpdateVoluntaryRequest", zap.Any("updateVoluntary", updateVoluntary))
	err = p.voluntaryFacade.UpdateVoluntary(ctx, voluntaryID, updateVoluntary)
	if err != nil {
		p.logger.Error("Error updating voluntary", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Voluntary updated successfully")
	return c.JSON(http.StatusOK, "Voluntary updated successfully")
}

// @Summary delete voluntary router
// @Description delete voluntary router
// @Tags Voluntary
// @Accept json
// @Param id path int true "id voluntary"
// @Produce json
// @Success 200 {string} string "Voluntary deleted successfully"
// @Failure 500 {object} error
// @Router /atados/v1/voluntary/{id} [delete]
func (p *voluntaryController) DeleteVoluntary(c echo.Context) error {
	p.logger.Info("Controller: Deleting voluntary")
	loggerUUID := c.Get("logger").(string)
	p.logger.Info("correlationID", zap.String("correlationID: ", loggerUUID))

	ctx := context.WithValue(context.Background(), "logger", loggerUUID)

	voluntaryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Param", zap.Uint64("voluntaryID", voluntaryID))
	err = p.voluntaryFacade.DeleteVoluntary(ctx, voluntaryID)
	if err != nil {
		p.logger.Error("Error deleting voluntary", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Voluntary deleted successfully")
	return c.JSON(http.StatusOK, "Voluntary deleted successfully")
}
