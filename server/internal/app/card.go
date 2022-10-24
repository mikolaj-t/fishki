package app

import (
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cardService struct {
	repo           core.CardRepository
	addHandlers    []core.CardObserverFunc
	removeHandlers []core.CardObserverFunc
}

func NewCardService(repo core.CardRepository) core.CardService {
	return &cardService{repo: repo}
}

func (c *cardService) Create(card *core.Card, deckId string) error {
	card.ID = primitive.NewObjectID().Hex()

	if err := c.repo.Create(card); err != nil {
		return err
	}

	for _, handler := range c.addHandlers {
		if err := handler(deckId, card.ID); err != nil {
			return err
		}
	}

	return nil
}

func (c *cardService) Get(id string) (*core.Card, error) {
	return c.repo.Get(id)
}

func (c *cardService) Update(card *core.Card) error {
	return c.repo.Update(card)
}

func (c *cardService) Delete(cardId string, deckId string) error {

	if deckId != "" {
		for _, handler := range c.removeHandlers {
			if err := handler(deckId, cardId); err != nil {
				return err
			}
		}
	}

	return c.repo.Delete(cardId)
}

func (c *cardService) AddCreatedCardObserver(observerFunc core.CardObserverFunc) {
	c.addHandlers = append(c.addHandlers, observerFunc)
}

func (c *cardService) AddDeletedCardObserver(observerFunc core.CardObserverFunc) {
	c.removeHandlers = append(c.removeHandlers, observerFunc)
}
