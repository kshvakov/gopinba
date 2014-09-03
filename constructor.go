package gopinba

import (
	"os"
)

func New(options *Options) *pinba {

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

	pinba := &pinba{
		hostname:   hostname,
		serverName: serverName,
		pinbaHost:  pinbaHost,
		pinbaPort:  pinbaPort,
	}

	return pinba
}
