package v1

import (
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/Jay-Chou118/mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorite(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	CreateFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&CreateFavoriteService); err != nil {
		res := CreateFavoriteService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}

func ListFavorite(c *gin.Context) {
	listFavoriteService := service.FavoriteService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listFavoriteService); err != nil {
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavoriteService); err != nil {
		res := deleteFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Infoln(err)
	}
}
