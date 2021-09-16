package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/models"
)

//return cancela el proceso.
func RegisterUser(w http.ResponseWriter, response *http.Request) {

	var user models.User

	err := json.NewDecoder(response.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email Required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "You must specify a password of at least 6 characters", 400)
		return
	}

	_, userFound, _ := database.UserExist(user.Email)
	if !userFound {
		http.Error(w, "There is already a registered user with that email", 400)
		return
	}

	_, status, err := database.AddUser(user)
	if err != nil {
		http.Error(w, "An error occurred while trying to register user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Failed to insert user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
