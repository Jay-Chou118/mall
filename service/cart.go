package service

import (
	"context"
	"github.com/Jay-Chou118/mall/dao"
	"github.com/Jay-Chou118/mall/model"
	"github.com/Jay-Chou118/mall/pkg/e"
	"github.com/Jay-Chou118/mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	//判断有没有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	CartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
	}
	err = CartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

func (service *CartService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(aId)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}

func (service *CartService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	cartId, _ := strconv.Atoi(aId)
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartByCartId(uint(cartId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
