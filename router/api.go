package router

import (
	"github.com/gin-gonic/gin"
	"go-vhr/controller"
	"go-vhr/middleware"
)

func InitRouter(engine *gin.Engine) {
	engine.POST("/doLogin", controller.Login)
	engine.GET("/logout", controller.Logout)
	engine.GET("/hr", controller.GetHrByName)
	engine.GET("/verifyCode", controller.GetCaptcha)

	system := engine.Group("/system")
	system.Use(middleware.JwtApiMiddleWare)
	system.GET("/hr", controller.GetHrList)
	system.GET("/config/menu", controller.GetMenu)
}
