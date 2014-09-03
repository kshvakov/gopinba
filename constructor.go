package gopinba

import (
	"fmt"
	"net"
	"os"
)

func New(options *Options) *Pinba {

	hostname := "unknown"
	serverName := "unknown"
	pinbaHost := "127.0.0.1"
	pinbaPort := 30002

	if value, err := os.Hostname(); err == nil {

		hostname = value
	}

	if options.ServerName != "" {

		serverName = options.ServerName
	}

	if options.PinbaHost != "" {

		pinbaHost = options.PinbaHost
	}

	if options.PinbaPort != 0 {

		pinbaPort = options.PinbaPort
	}

	pinba := &Pinba{
		hostname:   hostname,
		serverName: serverName,
		connected:  false,
	}

	if connect, err := net.Dial("udp", fmt.Sprintf("%s:%d", pinbaHost, pinbaPort)); err == nil {

		pinba.connect = connect
		pinba.connected = true
	}

	return pinba
}
