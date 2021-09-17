package database

import (
	"context"
	"time"

	"github.com/ArmandoSJ/wallet-digital/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCC(credit models.CreditCard, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("templatedb")
	col := db.Collection("creditCard")

	creditc := make(map[string]interface{})
	if credit.Credits > 0 {
		creditc["credits"] = credit.Credits
	}

	updtString := bson.M{
		"$set": creditc,
	}

	filtro := bson.M{"UserID": bson.M{"$eq": ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
