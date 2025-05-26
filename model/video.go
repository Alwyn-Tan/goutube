package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title       string
	Description string
}
