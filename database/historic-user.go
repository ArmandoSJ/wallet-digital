package database

import (
	"context"
	"time"

	"github.com/ArmandoSJ/wallet-digital/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPaymentDetail(ID string) ([]*models.PaymentDetail, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("templatedb")
	col := db.Collection("detallecompra")

	var result []*models.PaymentDetail

	condicion := bson.M{
		"usuaioid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	//opciones.SetSort(bson.D{{Key: "fechacompra", Value: -1}})
	//opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return result, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.PaymentDetail
		err := cursor.Decode(&registro)
		if err != nil {
			return result, false
		}
		result = append(result, &registro)
	}
	return result, true
}
