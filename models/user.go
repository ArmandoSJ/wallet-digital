package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB */
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   string             `bson:"UserID" json:"UserID,omitempty"` //user-name
	UserName string             `bson:"user-name" json:"user-name,omitempty"`
	PhoneID  string             `bson:"phoneID" json:"phoneID,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
}
