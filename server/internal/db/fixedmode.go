package db

import (
	"context"
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type fixedModeRepo struct {
	coll *mongo.Collection
}

func NewFixedModeRepository(db *mongo.Database) core.FixedModeRepository {
	return &fixedModeRepo{coll: db.Collection(reviewCollection)}
}

func (f fixedModeRepo) Get(reviewId string) (*core.FixedMode, error) {
	bytes, err := f.coll.FindOne(context.TODO(), bson.M{"_id": reviewId}).DecodeBytes()
	mode := &core.FixedMode{}
	err = bytes.Lookup("mode").Unmarshal(mode)

	return mode, err
}

func (f fixedModeRepo) Update(reviewID string, mode *core.FixedMode) error {
	review := core.Review{ModeID: core.Fixed, Mode: mode}
	_, err := f.coll.ReplaceOne(context.TODO(), bson.M{"_id": reviewID}, review)
	return err
}
