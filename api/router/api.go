package router

import (
	"github.com/gin-gonic/gin"
	"go-vhr/api/controller"
	"go-vhr/api/middleware"
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

	employee := engine.Group("/employee/basic")
	employee.Use(middleware.JwtApiMiddleWare)
	employee.GET("/", controller.GetEmployees)
	employee.DELETE("/:id", controller.DeleteEmployee)
	employee.GET("/nations", controller.GetAllNations)
	employee.GET("/positions", controller.GetAllPositions)
	employee.GET("/joblevels", controller.GetAllJobLevels)
	employee.GET("/politicsstatus", controller.GetAllPoliticsstatus)
	employee.GET("/deps", controller.GetAllDepartments)
}
