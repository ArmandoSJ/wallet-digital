package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ArmandoSJ/wallet-digital/database"
)

func PaymentDetail(w http.ResponseWriter, response *http.Request) {

	ID := response.URL.Query().Get("usuarioid")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el id del usuario", http.StatusBadRequest)
		return
	}

	if len(response.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(response.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := database.GetPaymentDetail(ID, pag)
	log.Print(respuesta)
	log.Print(correcto)
	if !correcto {
		http.Error(w, "Error al leer los servicios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
