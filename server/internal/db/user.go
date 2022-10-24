package db

import (
	"context"
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type userRepository struct {
	baseRepository[core.User]
}

func NewUserRepository(db *mongo.Database) core.UserRepository {
	return &userRepository{baseRepository[core.User]{db.Collection(userCollection)}}
}

func (u userRepository) Create(user *core.User) error {
	return u.baseRepository.Create(*user)
}

func (u userRepository) Get(id string) (*core.User, error) {
	user := &core.User{}
	err := u.baseRepository.Get(user, id)
	return user, err
}

func (u userRepository) GetByUsername(username string) (*core.User, error) {
	result := &core.User{}
	// case-insensitive search
	regex := bson.M{"$regex": primitive.Regex{Pattern: username, Options: "i"}}
	err := u.coll.FindOne(context.TODO(), bson.M{"username": regex}).Decode(result)
	return result, err
}

func (u userRepository) GetByEmail(email string) (*core.User, error) {
	result := &core.User{}
	// case-insensitive search
	regex := bson.M{"$regex": primitive.Regex{Pattern: email, Options: "i"}}
	err := u.coll.FindOne(context.TODO(), bson.M{"email": regex}).Decode(result)
	return result, err
}

func (u userRepository) Update(user *core.User) error {
	return u.baseRepository.Update(*user, user.ID)
}

func (u userRepository) AddDeck(id string, deckID string) error {
	_, err := u.coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"decks." + deckID: struct{}{}}})
	return err
}

func (u userRepository) RemoveDeck(id string, deckID string) error {
	_, err := u.coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$pull": bson.M{"decks": deckID}})
	return err
}

func (u userRepository) AddReview(id string, reviewID string) error {
	_, err := u.coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"reviews." + reviewID: struct{}{}}})
	return err
}

func (u userRepository) RemoveReview(id string, reviewID string) error {
	_, err := u.coll.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$pull": bson.M{"reviews": reviewID}})
	return err
}
