package dao

import (
	"context"
	"github.com/Jay-Chou118/mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByAid(aid uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", aid).First(&cart).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id = ?", uId).Find(&cart).Error
	return
}

func (dao *CartDao) UpdateCartByUserId(cId uint, cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ?", cId).Updates(&cart).Error
}

func (dao *CartDao) DeleteCartByCartId(cId, uId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id=?", cId, uId).Delete(&model.Cart{}).Error
}

func (dao *CartDao) UpdateCartNumById(cId uint, num uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Update("num", num).Error
}
