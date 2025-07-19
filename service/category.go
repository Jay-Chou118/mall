package service

import (
	"context"
	"github.com/Jay-Chou118/mall/dao"
	"github.com/Jay-Chou118/mall/pkg/e"
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/Jay-Chou118/mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))
}
