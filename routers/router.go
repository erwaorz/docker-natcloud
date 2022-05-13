package routers

import (
	"docker-api-service/docker"
	"docker-api-service/jwt"
	"github.com/gin-gonic/gin"
)
import setting "docker-api-service/setting"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.Use(jwt.JWT())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "连接成功",
		})
	})
	api := r.Group("/api")
	api.POST("/create", docker.CreateDocker)
	api.POST("/info", docker.Dockerinfo)
	api.POST("/delete", docker.DeleteDocker)
	api.POST("/addport", docker.AddPort)
	api.POST("/delport", docker.DelPort)
	api.POST("/pause", docker.PauseDocker)
	api.POST("/unpause", docker.UnPauseDocker)
	api.POST("/boot", docker.StartDocker)
	api.POST("/reboot", docker.RestarDocker)
	api.POST("/stop", docker.StopDocker)
	return r
}
