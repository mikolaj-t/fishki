package db

import (
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/mongo"
)

const cardsCollection = "cards"

type cardRepository struct {
	baseRepository[core.Card]
}

func NewCardRepository(db *mongo.Database) core.CardRepository {
	return &cardRepository{baseRepository[core.Card]{coll: db.Collection(cardsCollection)}}
}

func (c cardRepository) Create(card *core.Card) error {
	return c.baseRepository.Create(*card)
}
func (c cardRepository) Get(id string) (*core.Card, error) {
	card := &core.Card{}
	err := c.baseRepository.Get(card, id)
	return card, err
}

func (c cardRepository) Update(card *core.Card) error {
	return c.baseRepository.Update(*card, card.ID)
}
