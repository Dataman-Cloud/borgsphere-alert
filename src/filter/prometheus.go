package filter

import (
	"fmt"
	"strings"

	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	Fields     []string
	Type       string
	Help       string
	Name       string
	CounterVec map[string]*prometheus.CounterVec
}

func init() {
	ps := &Prometheus{
		CounterVec: make(map[string]*prometheus.CounterVec),
	}
	pc, ok := config.GetConfig().GetFilterByModule("prometheus")
	if !ok {
		return
	}

	err := config.ParseConfig(pc, ps)
	if err != nil {
		log.Fatalf("parse prometheus config to yaml error: %v", err)
		return
	}

	for _, v := range ps.Fields {
		fields := strings.Split(v, ",")
		ps.CounterVec[v] = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: ps.Name + "_" + strings.Join(fields, "_"),
				Help: ps.Help,
			},
			fields,
		)
		prometheus.MustRegister(ps.CounterVec[v])
	}

	RegistryFilterModule("prometheus", ps)
}

func (p *Prometheus) Read(msg map[string]interface{}) map[string]interface{} {

	for k, v := range p.CounterVec {
		fields := strings.Split(k, ",")
		if len(fields) == 0 {
			continue
		}

		var values []string
		for _, field := range fields {
			data, ok := msg[field]
			if !ok {
				break
			}
			values = append(values, fmt.Sprint(data))
		}

		if len(fields) == len(values) {
			v.WithLabelValues(values...).Inc()
		}
	}
	return msg
}
