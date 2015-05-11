package gopinba

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	ProtoMessage "github.com/kshvakov/gopinba/Pinba"
	"net"
	"os"
	"runtime"
	"time"
)

var memStats runtime.MemStats

// Pinba access object
type Pinba struct {
	hostname   string
	serverName string
	connect    net.Conn
	connected  bool
}

// New Pinba request object
func (pinba *Pinba) Request() *request {

	runtime.ReadMemStats(&memStats)

	request := &request{
		timeStart:   time.Now(),
		scriptName:  os.Args[0],
		memoryUsage: memStats.TotalAlloc,
		timers:      make([]*timer, 0, 10),
	}

	return request
}

// Flush request and send data to Pinba server
func (pinba *Pinba) Flush(request *request) error {

	if !pinba.connected {

		return fmt.Errorf("Could not connect to pinba server")
	}

	runtime.ReadMemStats(&memStats)

	req := ProtoMessage.Request{
		Hostname:        proto.String(pinba.hostname),
		ServerName:      proto.String(pinba.serverName),
		ScriptName:      proto.String(request.scriptName),
		RequestCount:    proto.Uint32(1),
		DocumentSize:    proto.Uint32(0),
		MemoryPeak:      proto.Uint32(0),
		MemoryFootprint: proto.Uint32(uint32(memStats.TotalAlloc - request.memoryUsage)),
		RequestTime:     proto.Float32(float32(time.Since(request.timeStart).Seconds())),
		RuUtime:         proto.Float32(0),
		RuStime:         proto.Float32(0),
		Schema:          proto.String(request.schema),
		Dictionary:      make([]string, 0, 20),
		TimerTagName:    make([]uint32, 0, 20),
		TimerTagValue:   make([]uint32, 0, 20),
	}

	for _, timer := range request.timers {

		req.TimerTagCount = append(req.TimerTagCount, uint32(len(*timer.tags)))
		req.TimerHitCount = append(req.TimerHitCount, 1)
		req.TimerValue = append(req.TimerValue, float32(timer.timeEnd.Sub(timer.timeStart).Seconds()))

		for tag, value := range *timer.tags {

			if pos, exists := inSlice(tag, req.Dictionary); exists {

				req.TimerTagName = append(req.TimerTagName, uint32(pos))

			} else {

				req.Dictionary = append(req.Dictionary, tag)

				req.TimerTagName = append(req.TimerTagName, uint32(len(req.Dictionary)-1))
			}

			if pos, exists := inSlice(value, req.Dictionary); exists {

				req.TimerTagValue = append(req.TimerTagValue, uint32(pos))

			} else {

				req.Dictionary = append(req.Dictionary, value)

				req.TimerTagValue = append(req.TimerTagValue, uint32(len(req.Dictionary)-1))
			}
		}
	}

	message, err := proto.Marshal(&req)

	if err != nil {

		return err
	}

	if _, err = pinba.connect.Write(message); err != nil {

		return err
	}

	return nil
}
