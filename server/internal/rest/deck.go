package rest

import (
	"fishki/pkg/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type deckHandler struct {
	service core.DeckService
}

func NewDeckHandler(service core.DeckService) core.DeckHandler {
	return &deckHandler{service}
}

func (d deckHandler) Setup(router *gin.Engine, authorized *gin.RouterGroup) {
	router.GET("/decks/get", d.GetDeck)
	authorized.POST("/decks/create", d.Create)
	authorized.POST("/decks/update", d.Update)
	authorized.POST("/decks/delete", d.Delete)
	authorized.POST("/decks/cards/add", d.AddCard)
	authorized.POST("/decks/card/remove", d.RemoveCard)
}

func (d deckHandler) GetDeck(ctx *gin.Context) {
	id := ctx.Query("id")

	deck, err := d.service.Get(id)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.IndentedJSON(http.StatusOK, deck)
}

func (d deckHandler) Create(ctx *gin.Context) {
	deck := core.Deck{}

	user, found := UserFromCtx(ctx)

	if !found {
		ctx.Status(http.StatusForbidden)
		return
	}

	err := ctx.BindJSON(&deck)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = d.service.CreateDeck(&deck, user.ID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.IndentedJSON(http.StatusOK, deck)
}

func (d deckHandler) Update(ctx *gin.Context) {
	deck := core.Deck{}

	err := ctx.BindJSON(&deck)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if !d.authorizedUser(ctx, deck.ID) {
		return
	}

	err = d.service.Update(&deck)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.IndentedJSON(http.StatusOK, deck)
}

func (d deckHandler) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	if !d.authorizedUser(ctx, id) {
		return
	}

	err := d.service.Delete(id)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

func (d deckHandler) AddCard(ctx *gin.Context) {
	deckId := ctx.Query("id")
	cardId := ctx.Query("card")

	if !d.authorizedUser(ctx, deckId) {
		return
	}

	err := d.service.AddCard(deckId, cardId)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

func (d deckHandler) RemoveCard(ctx *gin.Context) {
	deckId := ctx.Query("id")
	cardId := ctx.Query("card")

	err := d.service.RemoveCard(deckId, cardId)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

func (d deckHandler) authorizedUser(ctx *gin.Context, deckId string) bool {
	user, _ := UserFromCtx(ctx)
	if !user.DecksID.Has(deckId) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return false
	}
	return true
}
