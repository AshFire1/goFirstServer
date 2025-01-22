package auth

import (
	"strconv"
	"time"

	"github.com/AshFire1/config"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiratation := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiratation).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
