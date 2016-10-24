package filter

import (
	"encoding/json"
	"fmt"

	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type Json struct {
	Field string
}

func init() {
	j := new(Json)
	jconfig, ok := config.GetConfig().GetFilterByModule("json")
	if !ok {
		return
	}

	err := config.ParseConfig(jconfig, j)
	if err != nil {
		log.Fatalf("parse filter json config to yaml error: %v", err)
		return
	}

	RegistryFilterModule("json", j)
}

func (j *Json) Read(msg map[string]interface{}) map[string]interface{} {
	mj := make(map[string]interface{})
	if err := json.Unmarshal([]byte(fmt.Sprint(msg[j.Field])), &mj); err == nil {
		for k, v := range mj {
			msg[k] = v
		}
	} else {
		msg["tag"] = "jsonparsefailed"
	}

	return msg
}
