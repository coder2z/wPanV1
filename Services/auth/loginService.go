package authService

import (
	"wPan/v1/Models"
	R "wPan/v1/Response"
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
	if err := Models.DB.Where("email = ?", l.Email).First(user).Error; err != nil {
		return "", R.LOGIN_PASSWORD_ERROR, false
	}
	if user.Status == 0 {
		return "", R.LOGIN_USER_BAN, false
	}
	if user.CheckPassword(l.PassWord) == false {
		return "", R.LOGIN_PASSWORD_ERROR, false
	}
	token, ok := GetJwt(user)
	if !ok {
		return "", R.LOGIN_JWT_ERROR, false
	}
	return token, R.LOGIN_PASSWORD_ok, true

}
