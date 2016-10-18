package filter

import (
	"github.com/Dataman-Cloud/borgsphere-alert/src/grok"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type Grok struct {
	Match []string
	Gk    *grok.Grok
}

func init() {
	gs := new(Grok)
	err := config.ParseConfig(config.GetConfig().Filter["grok"], gs)
	if err != nil {
		log.Fatalf("parse grok config to yaml error: %v", err)
		return
	}

	gm, err := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	if err != nil {
		log.Fatalf("new grok error: %v", err)
		return
	}
	gs.Gk = gm

	RegistryFilterModule("grok", gs)
}

func (g *Grok) Read(msg string) map[string]string {
	source := map[string]string{"message": msg}
	for _, reg := range g.Match {
		if values, err := g.Gk.Parse(reg, msg); err == nil && len(values) > 0 {
			values["message"] = msg
			return values
		} else {
			source["tag"] = "grokparsefailed"
		}
	}

	return source
}
