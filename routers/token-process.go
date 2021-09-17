package routers

import (
	"errors"
	"strings"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email valor de Email usado en todos los EndPoints
var Email string

//UserID es el ID devuelto del modelo, que se usará en todos los EndPoints
var UserID string

//Proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, userFound, _ := database.UserExist(claims.Email)
		if userFound {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, userFound, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido, no encontrado")
	}
	return claims, false, string(""), err
}
