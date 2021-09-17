package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
)

func PaymentDetail(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el id del usuario", http.StatusBadRequest)
		return
	}
	respuesta, correcto := database.GetPaymentDetail(ID)
	if !correcto {
		http.Error(w, "Error al leer los servicios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
