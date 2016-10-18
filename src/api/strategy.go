package api

import (
	"github.com/gin-gonic/gin"
)

func (a *Alert) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
