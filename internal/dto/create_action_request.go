package dto

import (
	"atados/challenger/internal/model"
)

type CreateActionRequest struct {
	Name         string `json:"name" `
	Institution  string `json:"institution"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Address      string `json:"address"`
	Description  string `json:"description"`
} // @name CreateActionRequest

func (c *CreateActionRequest) ConvertToActionVO() *model.Action {
	return &model.Action{
		Name:         c.Name,
		Institution:  c.Institution,
		City:         c.City,
		Neighborhood: c.Neighborhood,
		Address:      c.Address,
		Description:  c.Description,
	}
}
