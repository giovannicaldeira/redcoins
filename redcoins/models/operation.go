package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type operation struct {
	gorm.Model
	operationID int `gorm:"primary_key"`
	qty         int
	date        time.Time
	userEmail   string
}
