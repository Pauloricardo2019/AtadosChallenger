package dto

import "atados/challenger/internal/model"

type CreateVoluntaryRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
} // @name CreateVoluntaryRequest

func (c *CreateVoluntaryRequest) ConvertToVoluntaryVO() *model.Voluntary {
	return &model.Voluntary{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Neighborhood: c.Neighborhood,
		City:         c.City,
	}
}
