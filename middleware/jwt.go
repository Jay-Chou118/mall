package middleware

import (
	"github.com/Jay-Chou118/mall/pkg/e"
	"github.com/Jay-Chou118/mall/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		toke := c.GetHeader("Authorization")
		if toke == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(toke)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}
		if code != e.Success {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
