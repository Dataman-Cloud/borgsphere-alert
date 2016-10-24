package api

import (
	"net/http"
	"time"

	logger "github.com/Dataman-Cloud/borgsphere-alert/src/utils/log"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Alert struct {
}

func Load(middleware ...gin.HandlerFunc) http.Handler {
	alert := new(Alert)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware...)
	r.Use(logger.Ginrus(log.StandardLogger(), time.RFC3339Nano, false))

	r.GET("/ping", alert.Create)
	r.POST("/ping", alert.Create)
	r.GET("/metrics", gin.WrapH(prometheus.Handler()))

	return r
}
