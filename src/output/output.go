package output

import (
//"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

//	log "github.com/Sirupsen/logrus"
)

type Output struct {
	OutputServer map[string]interface{}
}

var outputObject *Output

func RegistryOutputModule(module string, server interface{}) {
	if outputObject == nil {
		outputObject = &Output{
			OutputServer: make(map[string]interface{}),
		}
	}
	outputObject.OutputServer[module] = server
}

func GetOutput() *Output {
	return outputObject
}
