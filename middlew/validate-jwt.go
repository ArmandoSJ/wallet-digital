package middlew

import (
	"net/http"

	"github.com/ArmandoSJ/wallet-digital/routers"
)

//Permite validar el JWT que nos viene en la petición
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, vErr := routers.ProcesoToken(r.Header.Get("Authorization"))
		if vErr != nil {
			http.Error(w, "Error en el Token, "+vErr.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
