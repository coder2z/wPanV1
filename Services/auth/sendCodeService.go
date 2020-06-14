package authService

import (
	"fmt"
	Redis "wPan/v1/Cache"
	"wPan/v1/Config"
	"wPan/v1/Nsq"
	R "wPan/v1/Response"
	"wPan/v1/Utils"
)

type SendCodeService struct {
	Email string `json:"email" form:"email" binding:"required"`
}

func (s *SendCodeService) SendCode() (string, bool) {
	code := Utils.GetRandomString(6)
	if Redis.Exists("Register_" + s.Email) {
		return R.SENDCODE_EXISTS, false
	}
	_, err := Redis.Set("Register_"+s.Email, code, 60*5)
	if err != nil {
		fmt.Println(err.Error())
		return R.SENDCODE_ERROR, false
	}
	to := []string{
		s.Email,
	}
	err = Nsq.SendNSQ(to, "注册"+Config.ServerSetting.Name, "<h1>你的验证码为："+code+"</h1>")
	if err != nil {
		return err.Error(), false
	}
	return R.SENDCODE_OK, true
}
