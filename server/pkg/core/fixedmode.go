package core

import "github.com/gin-gonic/gin"

type FixedMode struct {
	Intervals []uint16         `json:"intervals" bson:"intervals,omitempty"`
	Levels    map[string]uint8 `json:"levels" bson:"levels,omitempty"`
	Dates     map[string]Date  `json:"dates" bson:"dates,omitempty"`
}

type FixedModeRepository interface {
	Get(reviewId string) (*FixedMode, error)
	Update(reviewID string, mode *FixedMode) error
}

type FixedModeService interface {
	ReviewModeService
	Create(deckId string, intervals []uint16) (*FixedMode, error)
}

type FixedModeHandler interface {
	CreateMode(ctx *gin.Context) (ReviewMode, error)
}
