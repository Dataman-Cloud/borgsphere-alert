package input

import (
	"fmt"
	"net/http"

	"github.com/Dataman-Cloud/borgsphere-alert/src/filter"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type HTTPInput struct {
	Host string
	Port int
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

func (h *HTTPInput) CreatServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filter.GetFilter().Msg <- "this is a test"
		fmt.Fprintf(w, "Hello astaxie!")
	})
	http.ListenAndServe(fmt.Sprintf("%s:%d", h.Host, h.Port), nil)
}
