package middleware

import (
	conf "auth-service/app/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	expireTime  = 86400
	prefixToken = "Bearer"
	roleKey     = "role"
	userKey     = "user"
	userNameKey = "userName"
)

type jwtToken struct {
}

type userAuth struct {
	userId   int
	userName string
}

var Auth *jwtToken
var (
	jwtSecretKey      = conf.Get().Auth.SecretKey // replace with your own secret key
	jwtTokenDuration  = time.Hour * 24            // customize token duration based on your needs
	jwtIssuer         = "auth-service"            // replace with your own issuer name
	jwtRefreshExpires = time.Hour * 24 * 30       // customize refresh token duration based on your needs
)

func init() {
	Auth = &jwtToken{}
}

func (*jwtToken) generateAccessToken(user *userAuth) (string, error) {
	claims := jwt.MapClaims{
		"userName": user.userName,
		"iss":      fmt.Sprintf("%s-%s", jwtIssuer, jwtSecretKey),
		"exp":      time.Now().Add(jwtTokenDuration).Unix(),
		"sub":      user.userName,
		"scopes":   []string{"user"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
func (*jwtToken) generateRefreshToken() (string, error) {
	claims := jwt.MapClaims{
		"iss":    fmt.Sprintf("%s-%s", jwtIssuer, jwtSecretKey),
		"exp":    time.Now().Add(jwtRefreshExpires).Unix(),
		"scopes": []string{"refresh_token"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
