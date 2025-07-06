package serializer

import (
	"github.com/Jay-Chou118/mall/conf"
	"github.com/Jay-Chou118/mall/model"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"user_name"`
	Nickname string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.UserName,
		Nickname: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
