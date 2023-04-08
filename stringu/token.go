package stringu

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jarvanstack/go_utils/logger"
)

// jwt 密钥生成
func TokenMarshal(mData map[string]interface{}, privateKey string, expSeconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	//过期
	claims["exp"] = time.Now().Unix() + expSeconds
	//字段
	for k, v := range mData {
		claims[k] = v
	}
	//token 编码
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	s, err := token.SignedString([]byte(privateKey))
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return s, nil
}
func TokenUnmarshal(token, privateKey string) (jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return claims, nil
}
