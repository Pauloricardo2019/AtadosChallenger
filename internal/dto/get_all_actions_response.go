package dto

import (
	"atados/challenger/internal/model"
	"time"
)

type ActionPagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
} // @name ActionPagination

type GetAllActionsResponse struct {
	Actions    []Action         `json:"actions"`
	Pagination ActionPagination `json:"pagination"`
} // @name GetAllActionsResponse

type Action struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name" `
	Institution  string    `json:"institution"`
	City         string    `json:"city"`
	Neighborhood string    `json:"neighborhood"`
	Address      string    `json:"address"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
} // @name Action

func (g *GetAllActionsResponse) ParseFromActionVO(actions []model.Action, limit, offset int, total int64) {
	g.Pagination.Limit = limit
	g.Pagination.Offset = offset
	g.Pagination.Total = total

	for _, action := range actions {
		g.Actions = append(g.Actions, Action{
			ID:           action.ID,
			Name:         action.Name,
			Institution:  action.Institution,
			City:         action.City,
			Neighborhood: action.Neighborhood,
			Address:      action.Address,
			Description:  action.Description,
			CreatedAt:    action.CreatedAt,
			UpdatedAt:    action.UpdatedAt,
		})
	}
}
