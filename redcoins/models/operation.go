package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Operation struct {
	gorm.Model
	Qty    int       `gorm:"not null" json:"qty"`
	Date   time.Time `gorm:"type:timestamp;not null" json:"date"`
	Value  float64   `gorm:"not null" json:"value"`
	UserID uint      `gorm:"not null" json:"user_id"`
}
