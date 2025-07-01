package model

import "grom.io/gorm"

type Category struct{
	gorm.Model
	Category string
}