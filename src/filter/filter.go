package filter

import (
	"reflect"
	"sync"

	//"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"
	"github.com/Dataman-Cloud/borgsphere-alert/src/output"
	//log "github.com/Sirupsen/logrus"
)

type Filter struct {
	RMutex  sync.Mutex
	CMutex  sync.Mutex
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
}

func GetFilter() *Filter {
	return f
}

func (f *Filter) Read() {
	for {
		select {
		case msg := <-f.Msg:
			var rv interface{}
			for _, v := range f.Filters {
				rv = reflect.ValueOf(v).
					MethodByName("Read").
					Call([]reflect.Value{reflect.ValueOf(msg)})[0].
					Interface()
			}

			for _, v := range output.GetOutput().OutputServer {
				reflect.ValueOf(v).MethodByName("Send").Call([]reflect.Value{reflect.ValueOf(rv)})
			}
		}
	}
}
