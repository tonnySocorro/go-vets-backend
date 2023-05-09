package jwt

import (
	"project/go-vets-backend/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("MiClaveMIW")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"birthDate": user.BirthDate,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myKey)
	return tokenString, err
}
