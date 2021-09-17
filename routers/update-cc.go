package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/models"
)

func UpdateCredits(w http.ResponseWriter, response *http.Request) {

	var cc models.CreditCard

	err := json.NewDecoder(response.Body).Decode(&cc)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = database.UpdateCC(cc, UserID)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
