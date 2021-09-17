package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB */
type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Servicename string             `bson:"Servicename" json:"Servicename,omitempty"` //user-name
	Description string             `bson:"description" json:"description,omitempty"`
	Price       int                `bson:"price" json:"price,omitempty"`
	Avalible    bool               `bson:"avalible" json:"avalible"`
	Category    string             `bson:"category" json:"category,omitempty"`
}
