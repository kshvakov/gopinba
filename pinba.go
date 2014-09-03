package gopinba

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/kshvakov/gopinba/Pinba"
	"net"
	"time"
)

type pinba struct {
	pinbaHost  string
	pinbaPort  int
	hostname   string
	serverName string
}

func (pinba *pinba) Request() *request {

	request := &request{
		timeStart:  time.Now(),
		scriptName: "unknown",
		//timers:     make([]*timer, 0, 10),
	}

	return request
}

func (pinba *pinba) Flush(request *request) error {
	// @todo: не открывать коннект каждый раз
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", pinba.pinbaHost, pinba.pinbaPort))

	if err != nil {

		return err
	}

	req := Pinba.Request{
		Hostname:     proto.String(pinba.hostname),
		ServerName:   proto.String(pinba.serverName),
		ScriptName:   proto.String(request.scriptName),
		RequestCount: proto.Uint32(1),
		DocumentSize: proto.Uint32(0),
		MemoryPeak:   proto.Uint32(0),
		RequestTime:  proto.Float32(float32(time.Since(request.timeStart).Seconds())),
		RuUtime:      proto.Float32(0),
		RuStime:      proto.Float32(0),
		Dictionary:   make([]string, 0),
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

	if _, err = conn.Write(message); err != nil {

		return err
	}

	return nil
}
