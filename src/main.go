package main

import (
	"net/http"
	"time"

	"github.com/Dataman-Cloud/borgsphere-alert/src/api"
	middle "github.com/Dataman-Cloud/borgsphere-alert/src/router/middleware"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.Infof("http server listen %s starting", config.GetConfig().AlertPort)

	server := &http.Server{
		Addr:           config.GetConfig().AlertPort,
		Handler:        api.Load(middle.Authenticate),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("http listen and serve error: %v", err)
	}
}
