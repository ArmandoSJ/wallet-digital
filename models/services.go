package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Servicename string             `bson:"Servicename" json:"Servicename,omitempty"` //user-name
	Description string             `bson:"description" json:"description,omitempty"`
	Price       float64            `bson:"price" json:"price,omitempty"`
	Available   bool               `bson:"available" json:"available"`
	Category    string             `bson:"category" json:"category,omitempty"`
}
