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

	if options.serverName != "" {

		serverName = options.serverName
	}

	if options.pinbaHost != "" {

		pinbaHost = options.pinbaHost
	}

	if options.pinbaPort != 0 {

		pinbaPort = options.pinbaPort
	}

	pinba := &pinba{
		hostname:   hostname,
		serverName: serverName,
		pinbaHost:  pinbaHost,
		pinbaPort:  pinbaPort,
	}

	return pinba
}
