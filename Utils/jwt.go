package Utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"wPan/v1/Config"
	"wPan/v1/Models"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

func GenerateToken(user *Models.User) (string, error) {
	claim := jwt.MapClaims{
		"username": user.UserName,
		"id":       user.ID,
		"email":    user.Email,
		"status":   user.Status,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + Config.JWTSetting.ExpiresAt*60*60,
		"iss":      Config.JWTSetting.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err := token.SignedString([]byte(Config.JWTSetting.JwtSecret))
	return tokens, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JWTSetting.JwtSecret), nil
	}
}

func ParseToken(tokens string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user.Id = int(claim["id"].(float64))
	user.Username = claim["username"].(string)
	user.Email = claim["email"].(string)
	user.Status = int(claim["status"].(float64))
	return
}
