package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID        string `gorm:"not null;unique"`
	Name      string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Phone     string `gorm:"not null;unique"`
	BirthDate string `gorm:"not null"`
}
