Start Krakend API gateway by running command: `krakend run -c krakend.json`

```html
krakend.json
{
  "version": 2,
  "extra_config": {
    "github_com/devopsfaith/krakend-cors": {
      "allow_origins": [
        "*"
      ],
      "expose_headers": [
        "Content-Length"
      ],
      "max_age": "12h",
      "allow_methods": [
        "GET",
        "HEAD",
        "POST"
      ]
    }
  },
  "timeout": "3000ms",
  "cache_ttl": "300s",
  "output_encoding": "json",
  "name": "hello",
  "endpoints": [
    {
      "endpoint": "/hello",
      "method": "GET",
      "output_encoding": "json",
      "extra_config": {},
      "backend": [
        {
          "url_pattern": "/",
          "encoding": "json",
          "sd": "static",
          "method": "GET",
          "extra_config": {},
          "host": [
            "http://localhost",
            "http://localhost:8000"
          ],
          "disable_host_sanitize": false
        }
      ]
    }
  ]
}
```

```html
gani@linux:~/Projects/golang$ ./krakend run -c hello-krakend.json
Parsing configuration file: hello-krakend.json
2020/10/19 16:03:43  ERROR: unable to create the gologging logger: getting the extra config for the krakend-gologging module
2020/10/19 16:03:43  ERROR: unable to create the GELF writer: getting the extra config for the krakend-gelf module
2020/10/19 16:03:43  INFO: Listening on port: 8080
2020/10/19 16:03:43  DEBUG: creating a new influxdb client
2020/10/19 16:03:43  DEBUG: no config for the influxdb client. Aborting
2020/10/19 16:03:43  WARNING: influxdb: unable to load custom config
2020/10/19 16:03:43  WARNING: opencensus: no extra config defined for the opencensus module
2020/10/19 16:03:43  WARNING: building the etcd client: unable to create the etcd client: no config
2020/10/19 16:03:43  DEBUG: no config for the bloomfilter
2020/10/19 16:03:43  WARNING: bloomFilter: no config for the bloomfilter
2020/10/19 16:03:43  WARNING: no config present for the httpsecure module
2020/10/19 16:03:43  DEBUG: lua: no extra config
2020/10/19 16:03:43  DEBUG: botdetector middleware:  no config defined for the module
2020/10/19 16:03:43  INFO: registering usage stats for cluster ID 'T0b0lqRZt7G9Vw9XvNrSHvqetnjMihcLPfPTlXl8QJI='
2020/10/19 16:03:43  DEBUG: AMQP: http://localhost: no amqp consumer defined
2020/10/19 16:03:43  DEBUG: AMQP: http://localhost: no amqp producer defined
2020/10/19 16:03:43  DEBUG: pubsub: subscriber (http://localhost): github.com/devopsfaith/krakend-pubsub/subscriber not found in the extra config
2020/10/19 16:03:43  DEBUG: pubsub: publisher (http://localhost): github.com/devopsfaith/krakend-pubsub/publisher not found in the extra config
2020/10/19 16:03:43  DEBUG: http-request-executor: no extra config for backend /
2020/10/19 16:03:43  DEBUG: CEL: no extra config detected for backend /
2020/10/19 16:03:43  DEBUG: lua: no extra config
2020/10/19 16:03:43  DEBUG: CEL: no extra config detected for pipe /hello
2020/10/19 16:03:43  DEBUG: lua: no extra config
2020/10/19 16:03:43  INFO: JOSE: signer disabled for the endpoint /hello
2020/10/19 16:03:43  DEBUG: lua: no extra config
2020/10/19 16:03:43  INFO: JOSE: validator disabled for the endpoint /hello
2020/10/19 16:03:43  DEBUG: botdetector:  no config defined for the module
2020/10/19 16:03:43  DEBUG: http-server-handler: no extra config
[GIN] 2020/10/19 - 16:03:51 | 500 |     564.785µs |       127.0.0.1 | GET      "/hello"
Error #01: Invalid status code
[GIN] 2020/10/19 - 16:13:18 | 500 |     822.255µs |       127.0.0.1 | GET      "/hello"
Error #01: invalid character '<' looking for beginning of value
[GIN] 2020/10/19 - 16:13:27 | 500 |     409.361µs |             ::1 | GET      "/hello"
Error #01: invalid character '<' looking for beginning of value
[GIN] 2020/10/19 - 16:13:28 | 200 |      780.72µs |             ::1 | GET      "/hello"
[GIN] 2020/10/19 - 16:15:21 | 200 |     1.02931ms |       127.0.0.1 | GET      "/hello"
```

**Backend**

```go
package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	http.ListenAndServe(":8000", c.Handler(handler))
}
```

**Test**

https://github.com/astaxie/bat

```go
gani@linux:~/Projects/golang$ ./bat http://localhost:8080/hello
GET /hello HTTP/1.1
Host: localhost:8080
Accept: application/json
Accept-Encoding: gzip, deflate
User-Agent: bat/0.1.0

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Vary: Origin
X-Krakend: Version 1.2.0
X-Krakend-Completed: true
Date: Mon, 19 Oct 2020 08:15:21 GMT
Content-Length: 17
Cache-Control: public, max-age=300


{
  "hello": "world"
}
```

