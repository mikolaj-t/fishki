package db

import (
	"context"
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const reviewCollection = "reviews"

type reviewRepository struct {
	baseRepository[core.Review]
}

func NewReviewRepository(db *mongo.Database) core.ReviewRepository {
	return &reviewRepository{baseRepository[core.Review]{db.Collection(reviewCollection)}}
}

func (r reviewRepository) Create(review *core.Review) error {
	return r.baseRepository.Create(*review)
}

func (r reviewRepository) Get(id string) (*core.Review, error) {
	res := r.coll.FindOne(context.TODO(), bson.M{"_id": id})
	bytes, err := res.DecodeBytes()

	if err != nil {
		return nil, err
	}

	rawValue := bytes.Lookup("modeID")
	modeID := core.ReviewModeID(rawValue.Int32())

	mode := core.ReviewModeFromID(modeID)
	err = bytes.Lookup("mode").Unmarshal(mode)

	if err != nil {
		return nil, err
	}

	review := &core.Review{}

	err = res.Decode(review)

	// Decode() can't handle the inner instance of ReviewMode struct and it throws an error,
	// but it is later unmarshalled, so we can disregard it.
	if strings.Contains(err.Error(), "core.ReviewMode") {
		err = nil
	}

	review.Mode = mode
	return review, err
}

func (r reviewRepository) Update(review *core.Review) error {
	return r.baseRepository.Update(*review, review.ID)
}
