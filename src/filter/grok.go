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
	gk, ok := config.GetConfig().GetFilterByModule("grok")
	if !ok {
		return
	}
	err := config.ParseConfig(gk, gs)
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

func (g *Grok) Read(msg map[string]interface{}) map[string]interface{} {
	for _, reg := range g.Match {
		if values, err := g.Gk.Parse(reg, msg["message"].(string)); err == nil && len(values) > 0 {
			for k, v := range values {
				msg[k] = v
			}
			return msg
		} else {
			msg["tag"] = "grokparsefailed"
		}
	}

	return msg
}
