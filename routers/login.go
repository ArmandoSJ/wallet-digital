package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/jwt"
	"github.com/ArmandoSJ/wallet-digital/models"
)

/*Login realiza el login */
func Login(w http.ResponseWriter, response *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	vErr := json.NewDecoder(response.Body).Decode(&user)
	if vErr != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+vErr.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "El correo del usuario es requerido ", 400)
		return
	}
	documento, existe := database.Ingresar(user.Email, user.Password)
	if !existe {
		http.Error(w, "Usuario y/o Contraseña inválidos ", 400)
		return
	}

	jwtKey, vErr := jwt.GenerateJWT(documento)
	if vErr != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+vErr.Error(), 400)
		return
	}

	resp := models.LoginToken{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//guardar el token en la cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
