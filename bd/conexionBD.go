package bd

import (
	"context" //espacio de memoria donde se pueden compartir diferentes cosas
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()                                                                                                                     // url que va a conectar
var clientOptions = options.Client().ApplyURI("mongodb+srv://TobiasBanno:Tb339059@redsocial.mvci4kr.mongodb.net/?retryWrites=true&w=majority") // setea URL de base de datos

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("No se pudo conectar a la BD: " + err.Error()) // usamos fun .Error ya que lo convierte el obj en string
		return client
	}
	err = client.Ping(context.TODO(), nil) //ve si la conexion está disponible

	if err != nil {
		log.Fatal("conexión no disponible en la BD: " + err.Error())
		return client
	}

	log.Println("Connect a BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
