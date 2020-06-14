package R

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//R.Ok(c, "自定义msg",data)
func Ok(c *gin.Context, msg string, data interface{}) {
	Response(c, SUCCESS, msg, data, http.StatusOK)
}

//R.Error(c, "自定义msg",data)
func Error(c *gin.Context, msg string, data interface{}) {
	Response(c, FAIL, msg, data, http.StatusOK)
}

//R.Response(c,1,"msg",data,200)
func Response(c *gin.Context, code int, msg string, data interface{}, status int) {
	c.Render(status, render.JSON{Data: Res{
		code,
		msg,
		data,
	}})
}
