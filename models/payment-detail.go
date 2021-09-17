package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario es el modelo de usuario de la base de MongoDB
type PaymentDetail struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ServiceID   string             `bson:"servicioid" json:"servicioid,omitempty"` //user-name
	UsuarioID   string             `bson:"usuarioid" json:"usuarioid,omitempty"`
	Cantidad    int                `bson:"cantidad" json:"cantidad,omitempty"`
	Total       float64            `bson:"total" json:"total,omitempty"`
	FechaCompra time.Time          `bson:"fechacompra" json:"fechacompra,omitempty"`
}
