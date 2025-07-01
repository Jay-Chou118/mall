package model

import "grom.io/gorm"

type Admin struct{
	gorm.Model
	UserName string
	PasswordDigest string
	Avatar string

}