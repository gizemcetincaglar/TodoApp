package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gizemcetincaglar/todo_app/models"
)

var jwtSecret = []byte("gizli-anahtar") 

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
