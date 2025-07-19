package dao

import (
	"context"
	"github.com/Jay-Chou118/mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(id, userId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id=?", id, userId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (order []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ?", uId).Find(&order).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id = ?", aId).Updates(&order).Error
}

func (dao *OrderDao) DeleteOrderByOrderId(aId, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id=?", aId, uId).Delete(&model.Order{}).Error
}

func (dao *OrderDao) ListOrderByCondition(conditon map[string]interface{}, page model.BasePage) (order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(conditon).Count(&total).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).Where(conditon).Offset((page.PageNum - 1) * (page.PageSize)).Find(&order).Error
	return
}
