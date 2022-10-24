package db

import (
	"context"
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const decksCollection = "decks"

type deckRepository struct {
	baseRepository[core.Deck]
}

func NewDeckRepository(db *mongo.Database) core.DeckRepository {
	return &deckRepository{baseRepository[core.Deck]{coll: db.Collection(decksCollection)}}
}

/*func (d deckRepository) Create(deck *core.Deck) error {
	_, err := d.coll.InsertOne(context.TODO(), deck)
	return err
}*/

func (d deckRepository) Get(id string) (*core.Deck, error) {
	deck := &core.Deck{}
	err := d.baseRepository.Get(deck, id)
	return deck, err
}

func (d deckRepository) Create(deck *core.Deck) error {
	return d.baseRepository.Create(*deck)
}

func (d deckRepository) Update(deck *core.Deck) error {
	return d.baseRepository.Update(*deck, deck.ID)
}

/*func (d deckRepository) Update(deck *core.Deck) error {
	_, err := d.coll.UpdateOne(context.TODO(), bson.M{"_id": deck.ID}, deck)
	return err
}

func (d deckRepository) Delete(id string) error {
	_, err := d.coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}*/

func (d deckRepository) AddCard(deckId string, cardId string) error {
	_, err := d.coll.UpdateOne(context.TODO(), bson.M{"_id": deckId}, bson.M{"$push": bson.M{"cards": cardId}})
	return err
}

func (d deckRepository) RemoveCard(deckId string, cardId string) error {
	_, err := d.coll.UpdateOne(context.TODO(), bson.M{"_id": deckId}, bson.M{"$pull": bson.M{"cards": cardId}})
	return err
}
