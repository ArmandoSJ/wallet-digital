package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB */
type CreditCard struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID        string             `bson:"UsuarioID" json:"UsuarioID,omitempty"`
	CardNumberHidden string             `bson:"cardNumberHidden" json:"cardNumberHidden,omitempty"`
	CardNumber       string             `bson:"cardNumber" json:"cardNumber,omitempty"`
	Brand            string             `bson:"brand" json:"brand,omitempty"`
	Cvv              bool               `bson:"cvv" json:"cvv"`
	ExpiracyDate     string             `bson:"ExpiracyDate" json:"ExpiracyDate,omitempty"`
	CardHolderName   string             `bson:"cardHolderName" json:"cardHolderName,omitempty"`
	Credits          int                `bson:"credits" json:"credits,omitempty"`
}
