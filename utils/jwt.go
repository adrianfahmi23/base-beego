package utils

import (
	"time"

	"github.com/beego/beego/v2/server/web"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GenerateJWT(user map[string]interface{}) (string, error) {
	signature_key := []byte(web.AppConfig.DefaultString("jwt_key", "my-key-secret"))
	signing_method := jwt.SigningMethodHS256

	claims := JwtClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    web.AppConfig.DefaultString("appname", "appname"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: user["Username"].(string),
		Email:    user["Email"].(string),
	}

	token := jwt.NewWithClaims(
		signing_method,
		claims,
	)

	signedToken, err := token.SignedString(signature_key)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Fungsi untuk memverifikasi token
func VerifyJWT(tokenString string) (*JwtClaim, error) {
	claims := &JwtClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(web.AppConfig.DefaultString("jwt_key", "my-key-secret")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
