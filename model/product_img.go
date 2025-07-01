package model

import "grom.io/gorm"

type ProductImg struct{
	gorm.Model
	ProductId uint gorm:"not null"
	ImgPath string
	
}