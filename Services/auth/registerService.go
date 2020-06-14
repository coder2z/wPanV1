package authService

import (
	Redis "wPan/v1/Cache"
	"wPan/v1/Models"
	R "wPan/v1/Response"
)

type RegisterService struct {
	UserName  string `json:"userName" form:"userName" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Tpassword string `form:"Tpassword" binding:"required,min=8,max=40"`
	Code      string `form:"code" binding:"required,len=6"`
}

func (u *RegisterService) valid() (bool, string) {
	if u.Password != u.Tpassword {
		return false, R.PASSWORD_T
	}
	redisCode, _ := Redis.Get("Register_" + u.Email)
	if string(redisCode) != `"`+u.Code+`"` {
		return false, R.EMAIL_CODE
	}
	_, _ = Redis.Delete("Register_" + u.Email)
	c := 0
	Models.DB.Model(&Models.User{}).Where("email = ?", u.Email).Or("user_name=?", u.UserName).Count(&c)
	if c > 0 {
		return false, R.REG_USER_EXIST
	}
	return true, ""
}

func (reg *RegisterService) Register() (string, bool) {
	if ok, err := reg.valid(); !ok {
		return err, false
	}
	user := Models.User{
		UserName: reg.UserName,
		Email:    reg.Email,
	}
	if err := user.SetPassword(reg.Password); err != nil {
		return R.REG_BCRYPT_ERROR, false
	}
	if err := Models.DB.Create(&user).Error; err != nil {
		return R.REG_ERROR, false
	}
	return R.REG_OK, true
}
