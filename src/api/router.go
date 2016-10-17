package api

import (
	"net/http"
	"time"

	"github.com/Dataman-Cloud/borgsphere-alert/src/dao"
	"github.com/Dataman-Cloud/borgsphere-alert/src/dao/mysql"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"
	logger "github.com/Dataman-Cloud/borgsphere-alert/src/utils/log"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Alert struct {
	Store dao.Store
}

func Load(middleware ...gin.HandlerFunc) http.Handler {
	alert := new(Alert)
	InitStore(alert)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware...)
	r.Use(logger.Ginrus(log.StandardLogger(), time.RFC3339Nano, false))

	r.GET("/ping", alert.Create)

	return r
}

func InitStore(alert *Alert) {
	switch config.GetConfig().AlertDbDriver {
	case "mysql":
		alert.Store = mysql.InitMysqlStore()
	default:
		log.Errorf("invalid db driver: %s", config.GetConfig().AlertDbDriver)
	}
}
