package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateTokenString(user types.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

func VerifyTokenString(tokenString string) (*jwt.Token, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, nil
}

func GetTokenHeader(c *gin.Context) string {
	tokenString := c.Request.Header.Get("Authorization")
	return tokenString
}

func GetTokenClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func VerifyTokenExp(claims jwt.MapClaims) bool {
	isValid := float64(time.Now().Unix()) < claims["exp"].(float64)
	return isValid
}

func VerifyTokenUser(claims jwt.MapClaims) (*types.User, error) {
	userId := int(claims["id"].(float64))
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("Could not get the user id from claims: %s", err.Error()))
	//}

	user, err := db.UserRepo.FindUser(userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find user from claims: %s", err.Error()))
	}

	return user, nil
}

func RequireAuth(c *gin.Context) (*types.User, error) {
	tokenString := GetTokenHeader(c)

	token, err := VerifyTokenString(tokenString)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Token could not be verified: %s", err.Error()))
	}

	if claims, ok := GetTokenClaims(token); ok && token.Valid {
		tokenValid := VerifyTokenExp(claims)
		if !tokenValid {
			return nil, errors.New("Token is expired")
		}

		loggedUser, err := VerifyTokenUser(claims)
		if err != nil {
			return nil, errors.New("User not found")
		}

		if loggedUser.ID == 0 {
			return nil, errors.New("User is invalid")
		}

		return loggedUser, nil
	} else {
		return nil, errors.New("An error happened during the authentication")
	}
}
