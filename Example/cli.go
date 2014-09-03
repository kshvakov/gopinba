package main

import (
	"github.com/kshvakov/gopinba"
)

func main() {

	pinba := gopinba.New(&gopinba.Options{PinbaHost: "127.0.0.1", PinbaPort: 30002})

	request := pinba.Request()

	for i := 0; i < 100; i++ {

		timer := request.TimerStart(&gopinba.Tags{"group": "cli"})

		// exec

		request.TimerStop(timer)
	}

	pinba.Flush(request)
}
