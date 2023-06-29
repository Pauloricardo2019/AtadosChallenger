package dto

import "atados/challenger/internal/model"

type UpdateActionRequest struct {
	Name         string `json:"name" `
	Institution  string `json:"institution"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Address      string `json:"address"`
	Description  string `json:"description"`
} // @name UpdateActionRequest

func (u *UpdateActionRequest) ConvertToActionVO() *model.Action {
	return &model.Action{
		Name:         u.Name,
		Institution:  u.Institution,
		City:         u.City,
		Neighborhood: u.Neighborhood,
		Address:      u.Address,
		Description:  u.Description,
	}
}
