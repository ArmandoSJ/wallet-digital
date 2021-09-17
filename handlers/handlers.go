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
	router.HandleFunc("/register", middlew.ValidateDataBase(routers.RegisterUser)).Methods("POST")
	router.HandleFunc("/login", middlew.ValidateDataBase(routers.Login)).Methods("POST")
	router.HandleFunc("/supplier/wallet/recharge", middlew.ValidateDataBase(middlew.ValidateJWT(routers.UpdateCredits))).Methods("PUT")
	router.HandleFunc("/supplier/wallet/historic", middlew.ValidateDataBase(middlew.ValidateJWT(routers.PaymentDetail))).Methods("GET")
	router.HandleFunc("/supplier/addservice", middlew.ValidateDataBase(middlew.ValidateJWT(routers.InsertService))).Methods("POST")
	router.HandleFunc("/liberet/services", middlew.ValidateDataBase(middlew.ValidateJWT(routers.GetServices))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
