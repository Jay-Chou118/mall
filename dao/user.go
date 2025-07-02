package dao

import (
	"context"
	"fmt"

	"github.com/Jay-Chou118/mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(da *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username 判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	fmt.Println(user, err)
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, false, err
	}

	return user, true, nil

}

func (dao *UserDao) CreateUser(user model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error

}
