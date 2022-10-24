package core

import "github.com/gin-gonic/gin"

type Handler interface {
	Setup(router *gin.Engine, authorized *gin.RouterGroup)
}
