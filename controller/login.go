package controller

import (
	"github.com/gin-gonic/gin"
	"go-vhr/models"
	"go-vhr/tools"
	"net/http"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	var form models.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
	}

	ok, hr := models.MatchPassword(form.UserName, form.Password)
	if !ok {
		c.JSON(http.StatusOK, models.ResponseError("用户名或密码错误,请重新输入", nil))
		return
	}
	hr.Password = ""

	m := make(map[string]interface{})
	m["hr"] = hr
	m["expire_time"] = strconv.FormatInt(time.Now().UnixMilli(), 10)
	token, err := tools.MakeToken(m)
	if err != nil {
		tools.Logger().Errorln("generate token err", err)
		c.JSON(http.StatusOK, models.ResponseError("登录失败,内部错误", nil))
		return
	}

	r := &models.LoginResult{
		Hr:    *hr,
		Token: token,
	}
	c.JSON(http.StatusOK, models.ResponseOK("登录成功", r))
}
