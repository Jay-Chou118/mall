package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/Jay-Chou118/mall/dao"
	"github.com/Jay-Chou118/mall/model"
	"github.com/Jay-Chou118/mall/pkg/e"
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/Jay-Chou118/mall/serializer"
)

type ProductService struct {
	Id            uint   `json:"id" from:"id"`
	Name          string `json:"name" from:"name"`
	CategoryId    uint   `json:"category_id" from:"category_id"`
	Title         string `json:"title" from:"title"`
	Info          string `json:"info" from:"info"`
	ImgPath       string `json:"img_path" from:"img_path"`
	Price         string `json:"price" from:"price"`
	DiscountPrice string `json:"discount_price" from:"discount_price"`
	OnSale        bool   `json:"on_sale" from:"on_sale"`
	Num           int    `json:"num" from:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)

	//以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Tittle:        service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}

	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}
