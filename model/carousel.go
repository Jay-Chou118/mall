package model

import "grom.io/gorm"

type Carousel struct{
	gorm.Model
	ImgPath string
	ProductId uint gorm:"not null"
}