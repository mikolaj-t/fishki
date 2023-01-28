package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type baseRepository[T any] struct {
	coll *mongo.Collection
}

func (r baseRepository[T]) Create(t T) error {
	_, err := r.coll.InsertOne(context.TODO(), t)
	return err
}

func (r baseRepository[T]) Get(result *T, id string) error {
	err := r.coll.FindOne(context.TODO(), bson.M{"_id": id}).Decode(result)
	return err
}

func (r baseRepository[T]) Update(t T, id string) error {
	_, err := r.coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": t})
	return err
}

func (r baseRepository[T]) Delete(id string) error {
	_, err := r.coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
