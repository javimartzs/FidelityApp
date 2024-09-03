package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null;unique"`
	Role     string `gorm:"not null"`
}
