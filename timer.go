package gopinba

import (
	"time"
)

// Tags for Pinba request timers
//
// Example:
//
//    Tag{"group": "goapp", "operation": "fetch"}
//
//    Tag{"group": "mysql", "operation": "connect", "from_server": "127.0.0.1", "to_server": "192.168.10.101"}
//
type Tags map[string]string

type timer struct {
	started   bool
	timeStart time.Time
	timeEnd   time.Time
	tags      *Tags
}
