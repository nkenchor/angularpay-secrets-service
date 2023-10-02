package persistence

import (
	"context"
	"fmt"
	"log"
	"secret/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	Secret    repository.SecretRepository
	dbService *mongo.Collection
}

func ConnectDB(dbhost, dbname string) (Repositories, error) {

	clientOptions := options.Client().ApplyURI(dbhost) // Connect to MongoDB

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return Repositories{}, err
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return Repositories{}, err
	}
	fmt.Println("Connected to MongoDB!")

	conn := db.Database(dbname)

	secretCollection := conn.Collection("secret")

	return Repositories{
		Secret:    NewSecretInfra(secretCollection),
		dbService: secretCollection,
	}, nil
}
