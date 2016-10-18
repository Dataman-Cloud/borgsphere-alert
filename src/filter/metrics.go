package filter

import (
	"sync"

	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type Metrics struct {
	Field   string `yaml:"field"`
	Counter map[string]interface{}
	RMutex  sync.Mutex
}

func init() {
	m := &Metrics{
		Counter: make(map[string]interface{}),
	}
	err := config.ParseConfig(config.GetConfig().Filter["metrics"], m)
	if err != nil {
		log.Fatalf("parse metrics config error: %v", err)
		return
	}

	RegistryFilterModule("metrics", m)
}

func (m *Metrics) Read(msg map[string]interface{}) map[string]interface{} {
	_, ok := msg[m.Field]
	if !ok {
		return msg
	}

	if m.Counter[m.Field] == nil {
		m.Counter[m.Field] = uint64(0)
	}
	m.Counter[m.Field] = m.Counter[m.Field].(uint64) + 1
	return m.Counter
}
