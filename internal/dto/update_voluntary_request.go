package dto

import "atados/challenger/internal/model"

type UpdateVoluntaryRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
} // @name UpdateVoluntaryRequest

func (u *UpdateVoluntaryRequest) ConvertToVoluntaryVO() *model.Voluntary {
	return &model.Voluntary{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Neighborhood: u.Neighborhood,
		City:         u.City,
	}
}
