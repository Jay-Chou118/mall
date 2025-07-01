package model

import "grom.io/gorm"

type Notice stuct{
	gorm.Model
	Text string gorm:"type:text"
	
}