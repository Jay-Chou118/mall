package serializer

import (
	"context"
	"github.com/Jay-Chou118/mall/conf"
	"github.com/Jay-Chou118/mall/dao"
	"github.com/Jay-Chou118/mall/model"
)

type Order struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreateAt     int64   `json:"create_at"`
	UpdateAt     int64   `json:"update_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          int     `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Adress       string  `json:"adress"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		Id:           order.ID,
		OrderNum:     order.OrderNum,
		CreateAt:     order.CreatedAt.Unix(),
		UpdateAt:     order.UpdatedAt.Unix(),
		UserId:       order.UserId,
		ProductId:    order.ProductId,
		BossId:       order.BossId,
		Num:          order.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Adress:       address.Address,
		Type:         order.Type,
		ProductName:  product.Name,
		ImgPath:      conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Money:        order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressId)
		if err != nil {
			continue
		}
		favorite := BuildOrder(item, product, address)
		orders = append(orders, favorite)
	}
	return orders
}
