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
	userName := c.PostForm("username")
	password := c.PostForm("password")

	if ok := models.MatchPassword(userName, password); !ok {
		c.JSON(http.StatusOK, models.ResponseError("用户名或密码错误,请重新输入", nil))
		return
	}

	m := make(map[string]interface{})
	m["username"] = userName
	m["expire_time"] = strconv.FormatInt(time.Now().UnixMilli(),10)
	token, err := tools.MakeToken(m)
	if err != nil {
		tools.Logger().Errorln("generate token err", err)
		c.JSON(http.StatusOK, models.ResponseError("登录失败,内部错误", nil))
		return
	}
	c.SetCookie("JSESSIONID", token, 60*60, "/", "", false, false)

	c.JSON(http.StatusOK, models.ResponseOK("登录成功", nil))
}
