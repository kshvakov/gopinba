package main

import (
	"fmt"
	"github.com/kshvakov/gopinba"
	"net/http"
)

var pinba *gopinba.Pinba

func handler(w http.ResponseWriter, r *http.Request) {

	request := pinba.Request()

	request.SetScriptName(r.URL.Path)

	for i := 0; i < 100; i++ {

		timer := request.TimerStart(&gopinba.Tags{"group": "web"})

		// exec

		request.TimerStop(timer)
	}

	pinba.Flush(request)

	fmt.Fprint(w, "handler")
}

func main() {

	pinba = gopinba.New(&gopinba.Options{PinbaHost: "127.0.0.1", PinbaPort: 30002, ServerName: "go-web.local"})

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
