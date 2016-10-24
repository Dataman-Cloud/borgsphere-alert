package input

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Dataman-Cloud/borgsphere-alert/src/filter"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type HTTPInput struct {
	Fields map[string]interface{} `yaml:"fields" json:"fields"`
	Host   string
	Port   int
}

func init() {
	h, ok := config.GetConfig().Input["http"]
	if !ok {
		return
	}

	httpInput := new(HTTPInput)
	if err := config.ParseConfig(h, httpInput); err != nil {
		log.Errorf("parse http input error: %v", err)
		return
	}
	RegisterInputModule("http", httpInput)
}

func (h *HTTPInput) CreatServer(data map[string]interface{}) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if h.Fields != nil {
			data = h.Fields
		}
		defer r.Body.Close()
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf("get http requset body error: %v", err)
			return
		}
		data["message"] = string(msg)
		strdata, _ := json.Marshal(data)
		filter.GetFilter().Msg <- string(strdata)
	})
	http.ListenAndServe(fmt.Sprintf("%s:%d", h.Host, h.Port), nil)
}
