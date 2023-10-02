package persistence

import (
	"context"
	"log"
	"reflect"
	"secret/domain/entity"
	"secret/sharedinfrastructure/helper"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SecretInfra struct {
	collection *mongo.Collection
}

func NewSecretInfra(collection *mongo.Collection) *SecretInfra {
	return &SecretInfra{collection}
}

func (r *SecretInfra) CreateSecret(c entity.SecretStruct) (interface{}, error) {

	var errorString string = ""
	errArray, validateErr := helper.ValidateStruct(c)
	if validateErr != nil {
		for _, perError := range errArray {
			errError := perError
			errAdd := errError.Error() + " "
			errorString += errAdd
		}
		log.Println("400", errorString, "entity.SecretStruct", validateErr)
		return errorString, validateErr
	}

	c.CreatedOn = time.Now().Format(time.RFC3339)
	// c.LastModified = ""
	c.Reference = uuid.New().String()
	c.ServiceReference = uuid.New().String()
	ref := c.Reference

	_, err := r.collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Println("400", "unable to insert data", "entity.SecretStruct", err)
		return "unable to insert data", err
	}

	return ref, nil
}

func (r *SecretInfra) UpdateSecret(ref string, c entity.SecretStruct) (interface{}, error) {
	c.ServiceReference = uuid.New().String()
	c.LastModified = time.Now().Format(time.RFC3339)
	var errorString string = ""
	errArray, validateErr := helper.ValidateStruct(c)
	if validateErr != nil {
		for _, perError := range errArray {
			errError := perError
			errAdd := errError.Error() + " "
			errorString += errAdd
		}
		log.Println("400", errorString, "entity.MemeStruct", validateErr)
		return errorString, validateErr
	}
	filter := bson.M{"reference": ref}
	update := bson.M{"$set": bson.M{
		"ServiceReference": c.ServiceReference,
		"Name":             c.Name,
		"Value":            c.Value,
		"LastModified":     c.LastModified,
	}}
	result, err := r.collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		log.Println("400", "the filter does not match any documents", "entity.SecretStruct", err)
	}

	log.Println("modified count: ", result.ModifiedCount)
	return c, nil
}

func (r *SecretInfra) DeleteSecret(ref string, c entity.SecretStruct) (interface{}, error) {
	filter := bson.M{"reference": ref}

	res, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("400", "tunable to delete document", "entity.SecretStruct", err)
		return "the filter does not match any documents", err
	}
	if res.DeletedCount == 0 {
		log.Println("DeleteOne() document not found:", res)
	} else {
		log.Println(reflect.TypeOf(res))
	}
	return "", nil
}

func (r *SecretInfra) GetSecretByRef(ref string, c entity.SecretStruct) (interface{}, error) {
	filter := bson.M{"reference": ref}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		log.Println("400", "the filter does not match any documents", "entity.SecretStruct", err)
		return "the filter does not match any documents", err
	}
	return c, nil
}

func (r *SecretInfra) GetAllSecret(c entity.SecretStruct, C []entity.SecretStruct) (interface{}, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		defer cursor.Close(context.Background())
		log.Println("400", "unable to find data", "entity.SecretStruct", err)
		return "unable to find data", err
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&c)
		if err != nil {
			log.Println("500", "unable to decode data", "entity.SecretStruct", err)
			return "unable to decode data", err
		}
		C = append(C, c)
	}
	return C, nil
}

func (r *SecretInfra) GetServiceSecretList(ref string, c entity.SecretStruct) (interface{}, error) {
	filter := bson.M{"service_reference": ref}

	err := r.collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		log.Println("400", "the filter does not match any documents", "entity.SecretStruct", err)
		return "the filter does not match any documents", err
	}
	return "", nil

}
