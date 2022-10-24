package core

import "github.com/gin-gonic/gin"

type Answer struct {
	ReviewID string       `json:"review"`
	ModeID   ReviewModeID `json:"modeID"`
	CardID   string       `json:"card"`
	Correct  bool         `json:"correct"`
	Duration float32      `json:"duration"`
}

type AnswerService interface {
	Handle(answer Answer) error
}

type AnswerHandler interface {
	Handler
	Submit(ctx *gin.Context)
}
