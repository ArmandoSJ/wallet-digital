package database

import (
	"context"
	"time"

	"github.com/ArmandoSJ/wallet-digital/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Funcion para dar de alta un usuario en la base de datos.
func AddUser(user models.User) (string, bool, error) {

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoCN.Database("templatedb")
	col := database.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(context, user)
	if err != nil {
		return "", false, err
	}

	//Objetenemos el objectid y lo convertimos  a string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
