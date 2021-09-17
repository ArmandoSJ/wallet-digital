package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/models"
)

func InsertService(w http.ResponseWriter, response *http.Request) {
	var service models.Service
	err := json.NewDecoder(response.Body).Decode(&service)

	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	/*
		registro := models.Service{
			Servicename: service.Servicename,
			Description: service.Description,
			Price:       service.Price,
			Available:   service.Available,
			Category:    service.Category,
		} */

	_, status, err := database.AddService(service)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el servicio", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
