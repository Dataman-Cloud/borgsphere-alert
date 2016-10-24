package output

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

type TCPOutput struct {
	Host string
	Port uint16
	Conn *net.TCPConn
}

func init() {
	tcpOutput := new(TCPOutput)
	if config.GetConfig().Output["tcp"] == nil {
		return
	}
	err := config.ParseConfig(config.GetConfig().Output["tcp"], tcpOutput)
	if err != nil {
		log.Fatalf("parse tcpoutput config error: %v", err)
		return
	}

	tcpOutput.Dial()

	RegistryOutputModule("tcp", tcpOutput)
}

func (t *TCPOutput) Send(msg interface{}) {
	data, _ := json.Marshal(msg)
	if _, err := t.Conn.Write(append(data, '\n')); err != nil {
		log.Errorf("tcp write error: %v", err)
	}
}

func (t *TCPOutput) Dial() {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port))
	if err != nil {
		log.Errorf("resolve tcp add error: %v", err)
		return
	}

	t.Conn, err = net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Errorf("dial tcp connect error: %v", err)
		return
	}
}
