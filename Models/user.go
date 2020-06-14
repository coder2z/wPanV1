package Models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"user_name" gorm:"type:varchar(100);unique_index"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	PassWord string `json:"pass_word" gorm:"type:varchar(100)"`
	Status   int    `json:"status" gorm:"DEFAULT:1"`
	gorm.Model
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.PassWord = string(bytes)
	return nil
}
