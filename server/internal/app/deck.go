package app

import (
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deckService struct {
	repo        core.DeckRepository
	cardService core.CardService
	userService core.UserService
}

func NewDeckService(repo core.DeckRepository, cardService core.CardService, userService core.UserService) core.DeckService {
	return &deckService{repo: repo, cardService: cardService, userService: userService}
}

func (d deckService) CreateDeck(deck *core.Deck, userID string) error {
	deck.ID = primitive.NewObjectID().Hex()
	deck.CardsID = nil
	if err := d.userService.AddDeck(userID, deck.ID); err != nil {
		return err
	}
	return d.repo.Create(deck)
}

func (d deckService) Get(id string) (*core.Deck, error) {
	return d.repo.Get(id)
}

func (d deckService) Update(deck *core.Deck) error {
	return d.repo.Update(deck)
}

func (d deckService) Delete(id string) error {
	deck, err := d.Get(id)

	if err != nil {
		return err
	}

	for _, id := range deck.CardsID {
		// deckId="" to delete without invoking RemoveCard from the deck
		if err := d.cardService.Delete(id, ""); err != nil {
			return err
		}
	}

	return d.repo.Delete(id)
}

func (d deckService) AddCard(deckId string, cardId string) error {
	return d.repo.AddCard(deckId, cardId)
}

func (d deckService) RemoveCard(deckId string, cardId string) error {
	return d.repo.RemoveCard(deckId, cardId)
}
