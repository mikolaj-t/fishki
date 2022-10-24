package rest

import (
	"fishki/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cardHandler struct {
	service core.CardService
}

func NewCardHandler(service core.CardService) core.CardHandler {
	return &cardHandler{service}
}

func (c cardHandler) Setup(router *gin.Engine, authorized *gin.RouterGroup) {
	router.GET("/cards/get", c.GetById)
	authorized.POST("/cards/create", c.Create)
	authorized.POST("/cards/update", c.Update)
	authorized.POST("/cards/delete", c.Delete)
}

func (c cardHandler) GetById(ctx *gin.Context) {
	id := ctx.Query("id")
	card, err := c.service.Get(id)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.IndentedJSON(http.StatusOK, card)
}

func (c cardHandler) Create(ctx *gin.Context) {
	deckid := ctx.Query("deck")
	card := core.Card{}

	err := ctx.BindJSON(&card)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.service.Create(&card, deckid)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, card)
}

func (c cardHandler) Update(ctx *gin.Context) {
	card := core.Card{}

	err := ctx.BindJSON(card)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(&card)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c cardHandler) Delete(ctx *gin.Context) {
	cardId := ctx.Query("id")
	deckId := ctx.Query("deck")

	err := c.service.Delete(cardId, deckId)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}
