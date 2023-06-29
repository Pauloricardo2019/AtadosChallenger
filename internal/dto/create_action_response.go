package dto

import "atados/challenger/internal/model"

type CreateActionVO struct {
	ID uint64 `json:"id"`
} // @name CreateActionVO

func (c *CreateActionVO) ParseFromActionVO(action *model.Action) {
	c.ID = action.ID
}
