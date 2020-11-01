package tools

import (
	"github.com/dgrijalva/jwt-go"
	"go-one-server/util/conf"
	"time"
)

var jwtSecret = func() string {
	if conf.Config.App.JwtSecret == "" {
		return "3bf6a2bf959f57a946139521a75acf0d"
	}
	return conf.Config.App.JwtSecret
}

type Claims struct {
	jwt.StandardClaims
}

func GenerateJWT(id string, issuer string, minute ...int) (string, error) {
	nowTime := time.Now()
	var expireTime time.Time
	if len(minute) > 0 {
		expireTime = nowTime.Add(time.Duration(minute[0]) * time.Minute)
	} else {
		expireTime = nowTime.Add(8 * time.Hour)
	}
	claims := Claims{
		jwt.StandardClaims{
			Audience:  MD5V(id),          // 受众
			ExpiresAt: expireTime.Unix(), // 失效时间
			Id:        MD5V(id),          // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    issuer,            // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "go-one-server",   // 主题
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret()))
	return token, err
}

func ParseJWT(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret()), nil
		})
	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	return nil, err
}
