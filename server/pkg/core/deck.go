package core

import "github.com/gin-gonic/gin"

type Deck struct {
	ID      string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"name,omitempty"`
	Public  bool     `json:"public" bson:"public,omitempty"`
	CardsID []string `json:"cards" bson:"cards,omitempty"`
}

type DeckRepository interface {
	Create(deck *Deck) error
	Get(id string) (*Deck, error)
	Update(deck *Deck) error
	Delete(id string) error

	AddCard(deckId string, cardId string) error
	RemoveCard(deckId string, cardId string) error
}

type DeckService interface {
	CreateDeck(deck *Deck, userID string) error
	Get(id string) (*Deck, error)
	Update(deck *Deck) error
	Delete(id string) error

	AddCard(deckId string, cardId string) error
	RemoveCard(deckId string, cardId string) error
}

type DeckHandler interface {
	Handler
	GetDeck(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	AddCard(ctx *gin.Context)
	RemoveCard(ctx *gin.Context)
}
