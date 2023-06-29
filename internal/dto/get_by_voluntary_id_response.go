package dto

import (
	"atados/challenger/internal/model"
	"time"
)

type GetVoluntaryByIDResponse struct {
	ID           uint64    `gorm:"primary_key;auto_increment;column:id"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Neighborhood string    `gorm:"column:neighborhood"`
	City         string    `gorm:"column:city"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
} // @name GetVoluntaryByIDResponse

func (g *GetVoluntaryByIDResponse) ParseFromVoluntaryVO(voluntary *model.Voluntary) {
	g.ID = voluntary.ID
	g.FirstName = voluntary.FirstName
	g.LastName = voluntary.LastName
	g.Neighborhood = voluntary.Neighborhood
	g.City = voluntary.City
	g.CreatedAt = voluntary.CreatedAt
	g.UpdatedAt = voluntary.UpdatedAt
}
