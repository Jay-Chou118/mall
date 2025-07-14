package v1

import (
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/Jay-Chou118/mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err != nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}

func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err != nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}

func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}
	if err := c.ShouldBind(&searchProductService); err != nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}
