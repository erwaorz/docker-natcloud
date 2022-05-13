package jwt

import (
	"docker-api-service/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc { //返回JSON
	return func(ctx *gin.Context) {
		if strings.Trim(ctx.GetHeader("apikey"), "") != "" && strings.Trim(ctx.GetHeader("apikey"), "") == strings.Trim(setting.APIKEY, "") {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "对接密匙错误",
			})
			ctx.Abort()
			return
		}
	}
}
