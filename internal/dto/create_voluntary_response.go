package dto

import "atados/challenger/internal/model"

type CreateVoluntaryVO struct {
	ID uint64 `json:"id"`
} // @name CreateVoluntaryVO

func (c *CreateVoluntaryVO) ParseFromVoluntaryVO(voluntary *model.Voluntary) {
	c.ID = voluntary.ID
}
