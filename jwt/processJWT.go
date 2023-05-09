package jwt

import (
	"errors"
	"project/go-vets-backend/db"
	"project/go-vets-backend/models"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var IDUsuario string

func ProcessJWT(tokenString string) error {
	myKey := []byte("MiClaveMIW")

	splitToken := strings.Split(tokenString, "Bearer")
	if len(splitToken) != 2 {
		return errors.New("Invalid Token format")
	}
	tokenString = strings.TrimSpace(splitToken[1])

	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found := db.CheckUserExists(claims.Email)
		if found {
			IDUsuario = claims.ID.Hex()
		}
		return nil
	}
	if !token.Valid {
		return errors.New("Invalid Token")
	}
	return nil
}
