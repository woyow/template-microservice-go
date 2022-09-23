package handler

import (
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}