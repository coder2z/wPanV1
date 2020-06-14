package authService

import (
	"wPan/v1/Models"
	R "wPan/v1/Response"
)

type ChangePWService struct {
	OldPassWord  string `json:"old_password" form:"old_password" binding:"required,min=8,max=40"`
	NewPassWord  string `json:"new_password" form:"new_password" binding:"required,min=8,max=40"`
	TNewPassWord string `json:"t_new_password" form:"t_new_password" binding:"required,min=8,max=40"`
}

func (u *ChangePWService) valid() (bool, string) {
	if u.NewPassWord != u.TNewPassWord {
		return false, R.PASSWORD_T
	}
	return true, ""
}

func (ch *ChangePWService) ChangePW(userId int) (string, bool) {
	if ok, msg := ch.valid(); !ok {
		return msg, false
	}
	user := new(Models.User)
	Models.DB.Where("id=?", userId).First(user)
	if user.CheckPassword(ch.OldPassWord) == false {
		return R.CHANGEPW_OLDPASSWORD_ERROR, false
	}
	if err := user.SetPassword(ch.NewPassWord); err != nil {
		return R.REG_BCRYPT_ERROR, false
	}
	if Models.DB.Save(user).Error != nil {
		return R.CHANGEPW_PASSWORD_ERROR, false
	}
	return R.CHANGEPW_PASSWORD_OK, true
}
