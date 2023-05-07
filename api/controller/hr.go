package controller

import (
	"github.com/gin-gonic/gin"
	"go-vhr/pkg/models"
)

func GetHrList(c *gin.Context) {
	hrList := models.ListHr()
	c.JSON(200, hrList)
}

func GetHrByName(c *gin.Context) {
	hr := models.GetHrByName("admin")
	c.JSON(200, hr)
}
