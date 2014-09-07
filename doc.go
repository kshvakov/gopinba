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
//  func simpleRequest(w http.ResponseWriter, r *http.Request) {
//
//    request := pinba.Request()
//
//    request.SetScriptName(r.URL.Path)
//
//    // exec
//
//    request.TimerStop(timer)
//
//    go pinba.Flush(request)
//
//    fmt.Fprintf(w, "simpleRequest %s", r.URL.Path)
//  }
//
//  func requestWithTimers(w http.ResponseWriter, r *http.Request) {
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
//    // exec
//
//    request.TimerStop(timer)
//
//    go pinba.Flush(request)
//
//    fmt.Fprintf(w, "requestWithTimers %s", r.URL.Path)
//  }
//
//  func main() {
//
//    pinba = gopinba.New(&gopinba.Options{PinbaHost: "127.0.0.2", PinbaPort: 30002, ServerName: "go-web.local"})
//
//    http.HandleFunc("/", simpleRequest)
//    http.HandleFunc("/timers/", requestWithTimers)
//    http.ListenAndServe(":8080", nil)
//  }
//
// Pinba reports
//
//  create table report_by_group_operation (
//    group_value varchar(64) default null,
//    operation_value varchar(64) default null,
//    req_count int(11) default null,
//    req_per_sec float default null,
//    hit_count int(11) default null,
//    hit_per_sec float default null,
//    timer_value float default null,
//    timer_median float default null,
//    index_value varchar(256) default null
//  ) engine=pinba default charset=latin1 comment='tag2_info:group,operation';
//
//  create table report_by_group_operation_from_to (
//    script_name varchar(128) default null,
//    group_value varchar(64) default null,
//    operation_value varchar(64) default null,
//    from_server_value varchar(64) default null,
//    to_server_value varchar(64) default null,
//    req_count int(11) default null,
//    req_per_sec float default null,
//    hit_count int(11) default null,
//    hit_per_sec float default null,
//    timer_value float default null,
//    timer_median float default null,
//    index_value varchar(256) default null
//  ) engine=pinba default charset=latin1 comment='tagN_report:group,operation,from_server,to_server';
//
//
package gopinba
