package routes

import (
	api "github.com/Jay-Chou118/mall/api/v1"
	"github.com/Jay-Chou118/mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		//商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("image/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user/update", api.UserUpdate)
			authed.POST("user/avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			//显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)

			//收藏夹操作
			authed.GET("favorites", api.ListFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)

			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.GET("address", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

		}
	}
	return r

}
