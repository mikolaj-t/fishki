package rest

import (
	"errors"
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type reviewHandler struct {
	service      *core.ReviewService
	modeHandlers map[core.ReviewModeID]*core.ReviewModeHandler
}

func NewReviewHandler(service *core.ReviewService) core.ReviewHandler {
	mhMap := make(map[core.ReviewModeID]*core.ReviewModeHandler)
	return &reviewHandler{service: service, modeHandlers: mhMap}
}

func (r reviewHandler) Setup(router *gin.Engine, authorized *gin.RouterGroup) {
	authorized.GET("/reviews/get", r.Get)
	authorized.POST("/reviews/create", r.Create)
	authorized.POST("/review/update", r.Update)
	authorized.POST("/review/delete", r.Delete)
}

func (r reviewHandler) Create(ctx *gin.Context) {
	modeIdInt, err := strconv.Atoi(ctx.PostForm("modeID"))
	deckId := ctx.PostForm("deckID")
	name := ctx.PostForm("name")

	user, found := UserFromCtx(ctx)

	if !found {
		ctx.Status(http.StatusForbidden)
		return
	}

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	modeId := core.ReviewModeID(modeIdInt)
	mode, err := (*r.modeHandlers[modeId]).CreateMode(ctx)

	if err != nil {
		return
	}
	review := core.Review{ModeID: modeId, Mode: mode, Name: name}

	if err := (*r.service).CreateReview(&review, deckId, user.ID); err != nil {
		ctx.Status(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	review.Mode = mode
	ctx.IndentedJSON(http.StatusOK, review)
}

func (r reviewHandler) Get(ctx *gin.Context) {
	id := ctx.Query("id")

	// TODO uncomment
	/*if !r.authorizedUser(ctx, id) {
		return
	}*/

	review, err := (*r.service).Get(id)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, review)
}

func (r reviewHandler) Update(ctx *gin.Context) {

	// TODO rework
	ctx.AbortWithStatus(http.StatusNotImplemented)

	/*	modeId, err := strconv.Atoi(ctx.PostForm("modeID"))

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		review := core.Review{ModeID: core.ReviewModeID(modeId)}

		if err := (*r.service).Update(&review); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.IndentedJSON(http.StatusOK, review)*/
}

func (r reviewHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (r reviewHandler) RegisterModeHandler(id core.ReviewModeID, handler core.ReviewModeHandler) {
	r.modeHandlers[id] = &handler
	//panic("implement me")
}

func (r reviewHandler) GetModeService(id core.ReviewModeID) (*core.ReviewModeHandler, error) {
	m, ok := r.modeHandlers[id]
	var err error = nil
	if !ok {
		err = errors.New("not found")
	}
	return m, err
	//panic("implement me")
}

func (r reviewHandler) authorizedUser(ctx *gin.Context, id string) bool {
	user, _ := UserFromCtx(ctx)
	if !user.ReviewsID.Has(id) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return false
	}
	return true
}
