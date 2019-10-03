package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type user struct {
	gorm.Model
	email      string `gorm:"primary_key"`
	password   string
	name       string
	birthDate  time.Time
	operations []operation `gorm:"ForeignKey:userEmail"`
}
