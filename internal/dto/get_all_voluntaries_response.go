package dto

import (
	"atados/challenger/internal/model"
	"time"
)

type VoluntaryPagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
} // @name VoluntaryPagination

type GetAllVoluntariesResponse struct {
	Voluntaries []Voluntary         `json:"voluntaries"`
	Pagination  VoluntaryPagination `json:"pagination"`
} // @name GetAllVoluntariesResponse

type Voluntary struct {
	ID           uint64    `json:"id"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Neighborhood string    `gorm:"column:neighborhood"`
	City         string    `gorm:"column:city"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
} // @name Voluntary

func (g *GetAllVoluntariesResponse) ParseFromVoluntaryVO(voluntaries []model.Voluntary, limit, offset int, total int64) {
	g.Pagination.Limit = limit
	g.Pagination.Offset = offset
	g.Pagination.Total = total

	for _, voluntary := range voluntaries {
		g.Voluntaries = append(g.Voluntaries, Voluntary{
			ID:           voluntary.ID,
			FirstName:    voluntary.FirstName,
			LastName:     voluntary.LastName,
			Neighborhood: voluntary.Neighborhood,
			City:         voluntary.City,
			CreatedAt:    voluntary.CreatedAt,
			UpdatedAt:    voluntary.UpdatedAt,
		})
	}
}
