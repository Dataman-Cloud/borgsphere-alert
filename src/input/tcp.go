package input

import (
	"encoding/json"
	"net"

	"github.com/Dataman-Cloud/borgsphere-alert/src/filter"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type TCPInput struct {
	Fileds map[string]string `yaml:"fields"`
	Host   string
	Port   int
	Server *net.TCPListener
}

func init() {
	t, ok := config.GetConfig().Input["tcp"]
	if !ok {
		return
	}

	tcpInput := new(TCPInput)
	err := config.ParseConfig(t, tcpInput)
	if err != nil {
		log.Errorf("get tcp input config error: %v", err)
		return
	}

	RegisterInputModule("tcp", tcpInput)
}

func (t *TCPInput) CreatServer(data map[string]interface{}) {
	var err error
	t.Server, err = net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(t.Host), t.Port, ""})
	if err != nil {
		log.Fatalf("create tcp listener error: %v", err)
	}
	defer t.Server.Close()

	for {
		conn, err := t.Server.AcceptTCP()
		if err != nil {
			log.Errorf("accept a tcp client error: %v", err)
			continue
		}

		defer conn.Close()
		go func() {
			msg := make([]byte, 1024)
			for {
				i, err := conn.Read(msg)
				if err != nil {
					break
				}
				data["message"] = string(msg[0:i])
				strdata, _ := json.Marshal(data)
				filter.GetFilter().Msg <- string(strdata)
			}
		}()
	}

}
