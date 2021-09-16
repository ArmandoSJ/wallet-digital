package main

import (
	"log"

	"github.com/ArmandoSJ/wallet-digital/database"
	"github.com/ArmandoSJ/wallet-digital/handlers"
)

func main() {
	if database.ValidateConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.RouterServices()
}
