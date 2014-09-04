package gopinba

import (
	"time"
)

type Tags map[string]string

type timer struct {
	started   bool
	timeStart time.Time
	timeEnd   time.Time
	tags      *Tags
}
