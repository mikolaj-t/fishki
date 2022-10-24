package rest

import (
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type answerHandler struct {
	service core.AnswerService
}

func NewAnswerHandler(service core.AnswerService) core.AnswerHandler {
	return &answerHandler{service}
}

func (a answerHandler) Setup(router *gin.Engine, authorized *gin.RouterGroup) {
	authorized.POST("/answer/submit", a.Submit)
}

func (a answerHandler) Submit(ctx *gin.Context) {
	answer := core.Answer{}

	if err := ctx.BindJSON(&answer); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		panic(err)
		return
	}

	user, found := UserFromCtx(ctx)

	if !found {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("??????", user)
		return
	}

	// todo bring back
	/*if !user.ReviewsID.Has(answer.ReviewID) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}*/

	if err := a.service.Handle(answer); err != nil {
		ctx.Status(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		panic(err)
		return
	}

	ctx.Status(http.StatusOK)
}
