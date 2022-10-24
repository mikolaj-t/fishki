package rest

import (
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type fixedModeHandler struct {
	service *core.FixedModeService
}

func NewFixedModeHandler(service *core.FixedModeService) core.FixedModeHandler {
	return &fixedModeHandler{service: service}
}

func (f fixedModeHandler) CreateMode(ctx *gin.Context) (core.ReviewMode, error) {
	deckId := ctx.PostForm("deckID")
	intervalsString := ctx.PostFormArray("intervals")
	var intervals []uint16

	for _, interval := range intervalsString {
		i, err := strconv.Atoi(interval)

		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}

		intervals = append(intervals, uint16(i))
	}

	mode, err := (*f.service).Create(deckId, intervals)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return mode, nil
}
