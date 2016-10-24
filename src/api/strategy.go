package api

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (a *Alert) Create(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	log.Info("this is a test haha: ", string(data), err)
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
