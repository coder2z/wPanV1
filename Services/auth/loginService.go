package authService

import (
	"fmt"
	"wPan/v1/Models"
	"wPan/v1/Utils"
)

type LoginService struct {
	Email    string `json:"email" form:"email" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func GetJwt(u *Models.User) (string, bool) {
	token, err := Utils.GenerateToken(u)
	if err != nil {
		return "", false
	}
	return token, true
}

func (l *LoginService) Login() (string, string, bool) {
	user := new(Models.User)
	if err := Models.DB.Where("email = ?", l.Email).First(&user).Error; err != nil {
		fmt.Println(err.Error())
		return "", "用户名或者密码错误", false
	}
	if user.CheckPassword(l.PassWord) == false {
		return "", "用户名或者密码错误", false
	}
	token, ok := GetJwt(user)
	if !ok {
		return "", "生成JWT失败", false
	}
	return token, "登录成功", true

}
