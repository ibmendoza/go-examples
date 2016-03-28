package main

import (
	"github.com/antonholmquist/jason"
	"log"
	"time"
)

func main() {

	exampleJSON :=
		`
{
  "server_id": "7EZN6QUDB6TWRAQI380DHE",
  "version": "0.7.9",
  "go": "go1.6",
  "host": "0.0.0.0",
  "auth_required": false,
  "ssl_required": false,
  "tls_required": false,
  "tls_verify": false,
  "max_connections": 65536,
  "ping_interval": 120000000000,
  "ping_max": 2,
  "http_port": 8222,
  "https_port": 0,
  "max_control_line": 1024,
  "max_pending_size": 10485760,
  "cluster_port": 0,
  "tls_timeout": 0.5,
  "port": 4222,
  "max_payload": 1048576,
  "start": "2016-03-26T14:07:32.7335156+08:00",
  "now": "2016-03-26T15:21:46.0870313+08:00",
  "uptime": "1h14m13s",
  "mem": 0,
  "cores": 4,
  "cpu": 0,
  "connections": 1,
  "total_connections": 12,
  "routes": 0,
  "remotes": 0,
  "in_msgs": 1355883,
  "out_msgs": 712347,
  "in_bytes": 433882560,
  "out_bytes": 227951040,
  "slow_consumers": 0,
  "http_req_stats": {
    "/": 2,
    "/connz": 1,
    "/routez": 1,
    "/subsz": 1,
    "/varz": 109
  }
}
`

	v, _ := jason.NewObjectFromBytes([]byte(exampleJSON))

	now, _ := v.GetString("now")
	in_msgs, _ := v.GetNumber("in_msgs")
	out_msgs, _ := v.GetNumber("out_msgs")
	in_bytes, _ := v.GetNumber("in_bytes")
	out_bytes, _ := v.GetNumber("out_bytes")

	p := log.Println

	p("now:", now)
	p("in_msgs:", in_msgs)
	p("out_msgs:", out_msgs)
	p("in_bytes:", in_bytes)
	p("out_bytes:", out_bytes)

	t1, _ := time.Parse(time.RFC3339, now)
	p(t1)

	p("Unix()")
	secs := t1.Unix()
	p(secs)

	t := time.Unix(secs, 0)
	p(t)

	//str := strconv.FormatInt(t1.Unix(), 10)
	//p(str)
}
