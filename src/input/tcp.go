package input

import (
	"net"

	"github.com/Dataman-Cloud/borgsphere-alert/src/filter"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type TCPInput struct {
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

func (t *TCPInput) CreatServer() {
	var err error
	t.Server, err = net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(t.Host), t.Port, ""})
	if err != nil {
		log.Fatalf("create tcp listener error: %v", err)
	}

	defer t.Server.Close()

	go filter.GetFilter().Read()

	for {
		conn, err := t.Server.AcceptTCP()
		if err != nil {
			log.Errorf("accept a tcp client error: %v", err)
			continue
		}

		defer conn.Close()
		go func() {
			data := make([]byte, 1024)
			for {
				i, err := conn.Read(data)
				if err != nil {
					break
				}
				filter.GetFilter().Msg <- string(data[0:i])
			}
		}()
	}

}
