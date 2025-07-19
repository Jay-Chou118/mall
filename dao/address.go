package dao

import (
	"context"
	"github.com/Jay-Chou118/mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByAid(aid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", aid).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUserId(uId uint) (address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", uId).Find(&address).Error
	return
}

func (dao *AddressDao) UpdateAddressByUserId(aId uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id = ?", aId).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddressByAddressId(aId, uId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id = ? AND user_id=?", aId, uId).Delete(&model.Address{}).Error
}
