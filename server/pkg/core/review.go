package core

import "github.com/gin-gonic/gin"

type ReviewModeID uint8

const (
	Fixed ReviewModeID = iota + 1
)

type Review struct {
	ID     string       `json:"id" bson:"_id,omitempty"`
	Name   string       `json:"name" bson:"name,omitempty"`
	ModeID ReviewModeID `json:"modeID" bson:"modeID,omitempty"`
	Mode   ReviewMode   `json:"mode" bson:"mode,omitempty"`
}

type ReviewRepository interface {
	Create(review *Review) error
	Get(id string) (*Review, error)
	Update(review *Review) error
	Delete(id string) error
}

type ReviewService interface {
	CreateReview(review *Review, deckID string, userID string) error
	Get(id string) (*Review, error)
	Update(review *Review) error
	Delete(id string) error
	RegisterModeService(id ReviewModeID, service ReviewModeService)
	GetModeService(id ReviewModeID) (*ReviewModeService, error)
}

type ReviewHandler interface {
	Handler
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterModeHandler(id ReviewModeID, handler ReviewModeHandler)
	GetModeService(id ReviewModeID) (*ReviewModeHandler, error)
}

func ReviewModeFromID(id ReviewModeID) ReviewMode {
	switch id {
	case Fixed:
		return &FixedMode{}
	default:
		return nil
	}
}

type ReviewMode interface {
}

type ReviewModeService interface {
	Create(deckID string, intervals []uint16) (*FixedMode, error)
	HandleAnswer(answer Answer) error
}

type ReviewModeHandler interface {
	CreateMode(ctx *gin.Context) (ReviewMode, error)
}
