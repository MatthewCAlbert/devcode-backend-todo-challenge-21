package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID              uint `gorm:"primaryKey;autoIncrement"`
	ActivityGroupID uint
	ActivityGroup   ActivityGroup
	Title           string `gorm:"varchar(255)" binding:"required"`
	IsActive        bool
	Priority        string `gorm:"varchar(15)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
