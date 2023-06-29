package dto

import (
	"atados/challenger/internal/model"
	"time"
)

type GetActionByIDResponse struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name" `
	Institution  string    `json:"institution"`
	City         string    `json:"city"`
	Neighborhood string    `json:"neighborhood"`
	Address      string    `json:"address"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
} // @name GetActionByIDResponse

func (g *GetActionByIDResponse) ParseFromActionVO(action *model.Action) {
	g.ID = action.ID
	g.Name = action.Name
	g.Institution = action.Institution
	g.City = action.City
	g.Neighborhood = action.Neighborhood
	g.Address = action.Address
	g.Description = action.Description
	g.CreatedAt = action.CreatedAt
	g.UpdatedAt = action.UpdatedAt
}
