package main

import (
	"net/http"
	"time"

	"github.com/Dataman-Cloud/borgsphere-alert/src/api"
	"github.com/Dataman-Cloud/borgsphere-alert/src/input"
	middle "github.com/Dataman-Cloud/borgsphere-alert/src/router/middleware"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"
	_ "github.com/Dataman-Cloud/borgsphere-alert/src/utils/log"

	log "github.com/Sirupsen/logrus"
)

func main() {
	go input.RunInput()

	log.Infof("http server listen %s starting", config.GetConfig().Addr)

	server := &http.Server{
		Addr:           config.GetConfig().Addr,
		Handler:        api.Load(middle.Authenticate),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("http listen and serve error: %v", err)
	}
}
