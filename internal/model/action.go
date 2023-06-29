package model

import "time"

type Action struct {
	ID           uint64    `json:"id" gorm:"primaryKey;column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Institution  string    `json:"institution" gorm:"column:institution"`
	City         string    `json:"city" gorm:"column:city"`
	Neighborhood string    `json:"neighborhood" gorm:"column:neighborhood"`
	Address      string    `json:"address" gorm:"column:address"`
	Description  string    `json:"description" gorm:"column:description"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Action) TableName() string {
	return "actions"
}
