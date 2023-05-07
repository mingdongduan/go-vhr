package controller

import (
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-vhr/pkg/models"
	"go-vhr/pkg/tools"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var store = base64Captcha.DefaultMemStore
var loginInfo = make(map[string]*models.Hr)

func Login(c *gin.Context) {
	var form models.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
	}

	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if !store.Verify(captchaId.(string), form.Code, true) {
			c.JSON(http.StatusOK, models.ResponseError("验证码错误", nil))
			return
		}
	} else {
		c.JSON(http.StatusOK, models.ResponseError("验证失败", nil))
		return
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

	loginInfo[token] = hr
	c.JSON(http.StatusOK, models.ResponseOK("登录成功", r))
}

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	delete(loginInfo, token)
	c.JSON(http.StatusOK, models.ResponseOK("注销成功", nil))
}

func GetCaptcha(c *gin.Context) {
	// 生成验证码
	driver := base64Captcha.NewDriverDigit(40, 120, 5, 0.4, 60)
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError("服务器内部错误", nil))
		return
	}

	session := sessions.Default(c)
	session.Set("captcha", id)
	_ = session.Save()

	i := strings.Index(b64s, ",")
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64s[i+1:]))
	c.Header("Content-Type", "image/png")
	io.Copy(c.Writer, decoder)
}
