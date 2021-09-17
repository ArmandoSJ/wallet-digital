package database

import (
	"context"
	"log"
	"time"

	"github.com/ArmandoSJ/wallet-digital/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPaymentDetail(ID string, pagina int64) ([]*models.PaymentDetail, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("templatedb")
	col := db.Collection("detallecompra")

	var result []*models.PaymentDetail

	condicion := bson.M{
		"usuarioid": ID,
	}

	log.Print(condicion)
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fechacompra", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	log.Print(cursor)
	if err != nil {
		log.Print("entro al error 1")
		return result, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.PaymentDetail
		log.Print(registro)
		err := cursor.Decode(&registro)
		if err != nil {
			log.Print("entro al error 2")
			return result, false
		}
		result = append(result, &registro)
	}
	return result, true
}
