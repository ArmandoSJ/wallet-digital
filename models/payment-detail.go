package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB */
type PaymentDetail struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ServiceID string             `bson:"ServiceID" json:"ServiceID,omitempty"` //user-name
	UsuarioID string             `bson:"UsuarioID" json:"UsuarioID,omitempty"`
	Cantidad  int                `bson:"Cantidad" json:"Cantidad,omitempty"`
	Total     bool               `bson:"total" json:"total"`
}
