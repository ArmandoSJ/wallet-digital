package jwt

import (
	"github.com/ArmandoSJ/wallet-digital/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Genera el encriptado con Json Web Token */
func GenerateJWT(user models.User) (string, error) {

	myKey := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":    user.Email,
		"nombre":   user.UserName,
		"userid":   user.UserID,
		"phone":    user.PhoneID,
		"Password": user.Password,
		"_id":      user.ID.Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
