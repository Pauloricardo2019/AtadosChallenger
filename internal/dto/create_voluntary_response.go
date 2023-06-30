package dto

import "atados/challenger/internal/model"

type CreateVoluntaryResponse struct {
	ID uint64 `json:"id"`
} // @name CreateVoluntaryResponse

func (c *CreateVoluntaryResponse) ParseFromVoluntaryVO(voluntary *model.Voluntary) {
	c.ID = voluntary.ID
}
