package filter

import (
	"encoding/json"
	"reflect"

	"github.com/Dataman-Cloud/borgsphere-alert/src/output"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"
	//log "github.com/Sirupsen/logrus"
)

type Filter struct {
	Msg     chan string
	Counter map[string]uint64
	Filters map[string]interface{}
}

var f *Filter

func RegistryFilterModule(module string, entity interface{}) {
	if f == nil {
		f = &Filter{
			Counter: make(map[string]uint64),
			Msg:     make(chan string),
			Filters: make(map[string]interface{}),
		}
	}
	f.Filters[module] = entity

	go f.Read()
}

func GetFilter() *Filter {
	return f
}

func (f *Filter) Read() {
	for {
		select {
		case msg := <-f.Msg:
			var rm map[string]interface{}
			json.Unmarshal([]byte(msg), &rm)

			var rv interface{}

			for _, v := range config.GetConfig().Filter {
				for k, _ := range v {
					rv = reflect.ValueOf(f.Filters[k]).
						MethodByName("Read").
						Call([]reflect.Value{reflect.ValueOf(rm)})[0].
						Interface()
				}
			}

			if output.GetOutput() != nil {
				for _, v := range output.GetOutput().OutputServer {
					reflect.ValueOf(v).MethodByName("Send").Call([]reflect.Value{reflect.ValueOf(rv)})
				}
			}
		}
	}
}
