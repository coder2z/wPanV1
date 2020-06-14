package Router

import (
	"github.com/gin-gonic/gin"
	"wPan/v1/Controllers"
	"wPan/v1/Middleware"
)

func ResRouter(r *gin.Engine) *gin.Engine {
	r = authRouter(r)
	return r
}

func authRouter(r *gin.Engine) *gin.Engine {

	authApi := r.Group("/api/auth")
	{
		authApi.POST("/login", Controllers.Login)
		authApi.POST("/register", Controllers.Register)
		authApi.GET("/info", Middleware.Auth(), Controllers.Info)
		authApi.POST("/changePW", Middleware.Auth(), Controllers.ChangePW)
		authApi.POST("/recoverPW", Controllers.RecoverPW)           //找回密码发送邮件
		authApi.POST("/recoverPWCheck", Controllers.RecoverPWCheck) //找回密码验证
		authApi.GET("/sendCode", Controllers.SendCode)
	}
	return r
}
