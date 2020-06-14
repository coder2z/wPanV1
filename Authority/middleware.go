package Authority

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	R "wPan/v1/Response"
	"wPan/v1/Utils"
)

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, _ := c.Get("userInfo")
		userInfo := value.(*Utils.UserInfo)
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := userInfo.Username
		if ok, _ := e.Enforce(sub, obj, act); ok {
			c.Next()
			return
		} else {
			R.Response(c, http.StatusUnauthorized, R.MSG401, nil, http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
}
