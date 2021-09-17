package database

import (
	"context"
	"log"
	"time"

	"github.com/ArmandoSJ/wallet-digital/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetServicesDB(status string) ([]*models.Service, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("templatedb")
	col := db.Collection("services")

	var result []*models.Service

	condicion := bson.M{
		"status": status,
	}

	opciones := options.Find()
	opciones.SetLimit(20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		return result, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.Service
		err := cursor.Decode(&registro)
		if err != nil {
			log.Print(err.Error())
			return result, false
		}
		result = append(result, &registro)
	}
	return result, true
}
