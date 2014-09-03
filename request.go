package gopinba

import (
	//"fmt"
	"runtime"
	"time"
)

var memStats runtime.MemStats

type request struct {
	timeStart  time.Time
	schema     string
	scriptName string
	timers     []*timer
}

func (request *request) SetSchema(schema string) {

	request.schema = schema
}

func (request *request) SetScriptName(scriptName string) {

	request.scriptName = scriptName
}

func (request *request) TimerStart(tags *Tags) *timer {

	runtime.ReadMemStats(&memStats)

	timer := &timer{
		started:     true,
		timeStart:   time.Now(),
		tags:        tags,
		memoryUsage: memStats.TotalAlloc,
	}

	request.timers = append(request.timers, timer)

	return timer
}

func (request *request) TimerStop(timer *timer) {

	runtime.ReadMemStats(&memStats)

	timer.started = false
	timer.timeEnd = time.Now()
	timer.memoryUsage = memStats.TotalAlloc - timer.memoryUsage
}
