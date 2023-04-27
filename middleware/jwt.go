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
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusOK, models.ResponseError("未登录，请先登录", nil))
		c.Abort()
		return
	}

	parse := tools.ParseToken(token)

	if parse["expire_time"] == nil || isNotExpire(parse["expire_time"].(string)) {
		c.JSON(http.StatusOK, models.ResponseError("登录超时，请重新登录", nil))
		c.Abort()
	}
	c.Set("hr", parse["hr"])
}

func isNotExpire(expire string) bool {
	parseInt, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		return time.Now().UnixMilli()-parseInt > expireTime
	}
	return false
}
