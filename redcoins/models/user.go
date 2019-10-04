package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"unique_index;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Name     string    `gorm:"not null" json:"name"`
	Birthday time.Time `gorm:"type:timestamp;not null" json:"birthday"`
}
