package controller

import (
	"github.com/gin-gonic/gin"
	"go-vhr/pkg/models"
	"net/http"
)

func GetMenu(c *gin.Context) {
	var id int64
	if value, exists := c.Get("hr"); exists {
		id = int64(value.(map[string]interface{})["id"].(float64))
	}

	menus := models.QueryMenu(id)

	c.JSON(http.StatusOK, menus)
}
