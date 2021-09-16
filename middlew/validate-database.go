package middlew

import (
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/database"
)

// Middlew que me permite conocer el estado de la BD
func ValidateDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if database.ValidateConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500) //cuando se pierde la conexion, muestra el error 500
			return
		}
		next.ServeHTTP(w, response)
	}
}
