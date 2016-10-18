package input

import (
	"reflect"

	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

var inputObject *Input

type Input struct {
	InputServer map[string]interface{}
}

func RegisterInputModule(module string, server interface{}) {
	if inputObject == nil {
		inputObject = &Input{
			InputServer: make(map[string]interface{}),
		}
	}

	inputObject.InputServer[module] = server
}

func RunInput() {
	for k, _ := range config.GetConfig().Input {
		obj, ok := inputObject.InputServer[k]
		if ok {
			go reflect.ValueOf(obj).MethodByName("CreatServer").Call([]reflect.Value{})
		} else {
			log.Errorf("cannot find input: %s", k)
		}
	}

}
