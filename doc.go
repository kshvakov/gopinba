// Package gopinba provides interface to push data to Pinba server
//
//
// Examples
//
// Example command line app
//
//  pinba := gopinba.New(&gopinba.Options{PinbaHost: "127.0.0.1", PinbaPort: 30002})
//
//  request := pinba.Request()
//
//  timer := request.TimerStart(&gopinba.Tag{
//      "group": "mysql",
//      "operation": "connect",
//      "from_server": "127.0.0.1",
//      "to_server": "192.168.10.101",
//  })
//
//  // connect to mysql server
//
//  request.TimerStop(timer)
//
//  timer = request.TimerStart(&gopinba.Tag{
//      "group": "mysql",
//      "operation": "select",
//      "from_server": "127.0.0.1",
//      "to_server": "192.168.10.101",
//  })
//
//  // select from mysql server
//
//  request.TimerStop(timer)
//
//  timer = request.TimerStart(&gopinba.Tag{
//    "group": "cli",
//    "operation": "fetch_data",
//  })
//
//  // fetch
//
//  request.TimerStop(timer)
//
//  pinba.Flush(request)
//
//
// Example web app
//
//  var pinba *gopinba.Pinba
//
//  func handler(w http.ResponseWriter, r *http.Request) {
//
//    request := pinba.Request()
//
//    request.SetScriptName(r.URL.Path)
//
//    timer := request.TimerStart(&gopinba.Tag{
//        "group": "mysql",
//        "operation": "connect",
//        "from_server": "127.0.0.1",
//        "to_server": "192.168.10.101",
//    })
//
//    // connect to mysql server
//
//    request.TimerStop(timer)
//
//    timer = request.TimerStart(&gopinba.Tag{
//        "group": "mysql",
//        "operation": "select",
//        "from_server": "127.0.0.1",
//        "to_server": "192.168.10.101",
//    })
//
//    // select from mysql server
//
//    request.TimerStop(timer)
//
//    timer = request.TimerStart(&gopinba.Tag{
//        "group": "web",
//        "operation": "fetch_data",
//    })
//
//    // fetch
//
//    request.TimerStop(timer)
//
//    go pinba.Flush(request)
//
//    fmt.Fprint(w, "handler")
//  }
//
//  func main() {
//
//    pinba = gopinba.New(&gopinba.Options{PinbaHost: "127.0.0.2", PinbaPort: 30002, ServerName: "go-web.local"})
//
//    http.HandleFunc("/", handler)
//    http.ListenAndServe(":8080", nil)
//  }
//
package gopinba
