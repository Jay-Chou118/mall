package serializer

import (
	"github.com/Jay-Chou118/mall/model"
	"github.com/Jay-Chou118/mall/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: util.Encrypt.AesDecoding(item.Money),
	}
}
