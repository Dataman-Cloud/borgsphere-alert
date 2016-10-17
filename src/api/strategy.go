package api

import (
	"github.com/Dataman-Cloud/borgsphere-alert/src/model"

	"github.com/gin-gonic/gin"
)

func (a *Alert) Create(ctx *gin.Context) {
	a.Store.CreateStrategy(&model.Strategy{})
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
