package v1

import (
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/Jay-Chou118/mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err != nil {
		res := orderPay.PayDown(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}

}
