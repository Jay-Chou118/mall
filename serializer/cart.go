package serializer

import (
	"context"
	"github.com/Jay-Chou118/mall/conf"
	"github.com/Jay-Chou118/mall/dao"
	"github.com/Jay-Chou118/mall/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	Name          string `json:"name"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
