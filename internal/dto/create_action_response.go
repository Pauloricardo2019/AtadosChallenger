package dto

import "atados/challenger/internal/model"

type CreateActionResponse struct {
	ID uint64 `json:"id"`
} // @name CreateActionResponse

func (c *CreateActionResponse) ParseFromActionVO(action *model.Action) {
	c.ID = action.ID
}
