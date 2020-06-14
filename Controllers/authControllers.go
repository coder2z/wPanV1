package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	R "wPan/v1/Response"
	authService "wPan/v1/Services/auth"
	"wPan/v1/Utils"
)

func Login(c *gin.Context) {
	var loginService authService.LoginService
	if err := c.ShouldBind(&loginService); err == nil {
		if data, msg, ok := loginService.Login(); ok {
			R.Ok(c, msg, data)
		} else {
			R.Error(c, msg, data)
		}
	} else {
		R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
	}
	return
}

func Register(c *gin.Context) {
	var regService authService.RegisterService
	if err := c.ShouldBind(&regService); err == nil {
		if msg, ok := regService.Register(); ok {
			R.Ok(c, msg, nil)
		} else {
			R.Error(c, msg, nil)
		}
	} else {
		R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
	}
	return
}

func Info(c *gin.Context) {
	value, _ := c.Get("userInfo")
	userInfo := value.(*Utils.UserInfo)
	R.Ok(c, R.SUCCESSMSG, userInfo)
}

func ChangePW(c *gin.Context) {
	value, _ := c.Get("userInfo")
	userInfo := value.(*Utils.UserInfo)
	var regService authService.ChangePWService
	if err := c.ShouldBind(&regService); err == nil {
		if msg, ok := regService.ChangePW(userInfo.Id); ok {
			R.Ok(c, msg, nil)
		} else {
			R.Error(c, msg, nil)
		}
	} else {
		R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
	}
	return
}

func RecoverPW(c *gin.Context) {
}

func RecoverPWCheck(c *gin.Context) {

}

func SendCode(c *gin.Context) {
	var sendCService authService.SendCodeService
	if err := c.ShouldBind(&sendCService); err == nil {
		if msg, ok := sendCService.SendCode(); ok {
			R.Ok(c, msg, nil)
		} else {
			R.Error(c, msg, nil)
		}
	} else {
		R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
	}
	return
}
