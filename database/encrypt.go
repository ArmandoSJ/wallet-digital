package database

import "golang.org/x/crypto/bcrypt"

//Funcion que me permite encriptar la password recibida
func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
