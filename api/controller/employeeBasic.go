package controller

import (
	"github.com/gin-gonic/gin"
	"go-vhr/pkg/models"
	"net/http"
	"strconv"
)

func GetEmployees(c *gin.Context) {
	var query models.QueryEmployee
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
		return
	}

	if query.Size == 0 {
		query.Size = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	err, responsePage := models.GetEmployees(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
		return
	}

	c.JSON(http.StatusOK, responsePage)
}

func GetAllNations(c *gin.Context) {
	err, nations := models.GetAllNations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
		return
	}

	c.JSON(http.StatusOK, nations)
}

func GetAllPositions(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllPositions())
}

func GetAllJobLevels(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllJobLevels())
}

func GetAllPoliticsstatus(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllPoliticsstatus())
}

func GetAllDepartments(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllDepartments(-1))
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	empid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
		return
	}
	models.DeleteEmployee(empid)
	c.JSON(http.StatusOK, models.ResponseOK("删除成功", nil))
}
