package main

import (
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"
	_ "github.com/Dataman-Cloud/borgsphere-alert/src/utils/log"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.Info("vim-go", config.GetConfig().AlertPort)
}
