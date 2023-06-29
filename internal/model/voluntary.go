package model

import "time"

type Voluntary struct {
	ID           uint64    `gorm:"primary_key;auto_increment;column:id"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	Neighborhood string    `gorm:"column:neighborhood"`
	City         string    `gorm:"column:city"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
}

func (Voluntary) TableName() string {
	return "voluntaries"
}
