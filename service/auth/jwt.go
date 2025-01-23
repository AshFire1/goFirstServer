package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AshFire1/config"
	"github.com/AshFire1/types"
	"github.com/AshFire1/utils"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	ContextKeyUserID = contextKey("userID")
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

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from user request
		token := getTokenFromRequest(r)

		//validate jwt
		tokenValid, err := validateToken(token)
		if err != nil {
			log.Printf("failure to validate token: %v", err)
			permissionDenied(w)
			return
		}
		if !tokenValid.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}
		//if valid fetch user id from token
		claims := tokenValid.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)
		userID, _ := strconv.Atoi(str)
		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failure to get user by id: %v", err)
			permissionDenied(w)
			return
		}
		//set context "userID" to user id
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKeyUserID, u.ID)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}
func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	return ""

}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(config.Envs.JWT_SECRET), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))

}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(ContextKeyUserID).(int)
	if !ok {
		return -1
	}
	return userID
}
