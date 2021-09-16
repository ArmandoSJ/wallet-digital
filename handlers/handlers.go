package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ArmandoSJ/wallet-digital/middlew"
	"github.com/ArmandoSJ/wallet-digital/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RouterServices() {
	router := mux.NewRouter()
	router.HandleFunc("/login", middlew.ValidateDataBase(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
