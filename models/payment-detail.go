package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentDetail struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID   string             `bson:"usuarioid" json:"usuarioid,omitempty"`
	ServiceID   string             `bson:"servicioid" json:"servicioid,omitempty"`
	Cantidad    int32              `bson:"cantidad" json:"cantidad,omitempty"`
	Total       float64            `bson:"total" json:"total,omitempty"`
	FechaCompra time.Time          `bson:"fechacompra" json:"fechacompra,omitempty"`
}
