package model

import (
	"time"

	"gorm.io/gorm"
)

type ActivityGroup struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"varchar(255)" binding:"required"`
	Email     string `gorm:"varchar(255)"`
	Todos     []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
