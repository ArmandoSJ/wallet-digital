package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
)

func GetServices(w http.ResponseWriter, response *http.Request) {

	if len(response.URL.Query().Get("status")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	status := response.URL.Query().Get("status")

	respuesta, correcto := database.GetServicesDB(status)
	if !correcto {
		http.Error(w, "Error al leer los servicios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
