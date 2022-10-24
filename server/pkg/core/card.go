package core

import "github.com/gin-gonic/gin"

type Card struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Prompt string `json:"prompt,omitempty" bson:"prompt,omitempty"`
	Answer string `json:"answer,omitempty" bson:"answer,omitempty"`
}

type CardRepository interface {
	Create(card *Card) error
	Get(id string) (*Card, error)
	Update(card *Card) error
	Delete(id string) error
}

type CardObserverFunc func(string, string) error

type CardService interface {
	Create(card *Card, deckId string) error
	Get(id string) (*Card, error)
	Update(card *Card) error
	Delete(cardId string, deckId string) error
	AddCreatedCardObserver(observerFunc CardObserverFunc)
	AddDeletedCardObserver(observerFunc CardObserverFunc)
}

type CardHandler interface {
	Handler
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
