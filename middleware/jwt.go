package middleware

import (
	"github.com/gin-gonic/gin"
	"go-vhr/models"
	"go-vhr/tools"
	"net/http"
	"strconv"
	"time"
)

const expireTime = 60 * 60 * 1000

func JwtApiMiddleWare(c *gin.Context) {
	if c.Request.URL.Path == "/doLogin" {
		return
	}
	cookie := tools.GetCookie(c.Request, "JSESSIONID")
	if len(cookie) == 0 {
		c.JSON(http.StatusOK, models.ResponseError("未登录，请先登录", nil))
		c.Abort()
		return
	}

	token := tools.ParseToken(cookie)

	if token["expire_time"] == nil || isNotExpire(token["expire_time"].(string)) {
		c.JSON(http.StatusOK, models.ResponseError("登录超时，请重新登录" , nil))
		c.Abort()
	}
}

func isNotExpire(expire string) bool {
	parseInt, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		return time.Now().UnixMilli() - parseInt > expireTime
	}
	return false
}
