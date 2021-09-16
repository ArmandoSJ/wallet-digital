package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD <password>*/
var MongoCN = DataBaseConn()
var optionsClient = options.Client().ApplyURI("mongodb+srv://asalazarj:Temporal$321@templatedb.b8lqq.mongodb.net/templatedb?retryWrites=true&w=majority")

/*Función que me permite conectar la BD */
func DataBaseConn() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), optionsClient)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*Ping a la base de datos, si hay una conexion regresa 1 y si no hay una conexion manda un 0 */
func ValidateConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
