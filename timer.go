package gopinba

import (
	"time"
)

type Tags map[string]string

type timer struct {
	started     bool
	timeStart   time.Time
	timeEnd     time.Time
	memoryUsage uint64
	tags        *Tags
}
