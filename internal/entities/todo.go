package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text   string
	Author string
}

