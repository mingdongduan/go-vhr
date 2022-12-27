package router

import (
	"github.com/gin-gonic/gin"
	"go-vhr/controller"
	"go-vhr/middleware"
)

func InitRouter(engine *gin.Engine) {
	engine.Use(middleware.JwtApiMiddleWare)
	engine.POST("/doLogin", controller.Login)
	engine.GET("/hr", controller.GetHrByName)

	group := engine.Group("/system/hr")
	group.GET("/", controller.GetHrList)
}
