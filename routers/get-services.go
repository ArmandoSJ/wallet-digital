package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetServices(w http.ResponseWriter, response *http.Request) {

	if len(response.URL.Query().Get("status")) < 1 {
		registro := bson.M{
			"status": http.StatusBadRequest,
			"error":  "Debe enviar el estado",
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(registro)
		//http.Error(w, registro , http.StatusBadRequest)
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
