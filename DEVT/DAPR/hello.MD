```go
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Running at port 3000...")

	http.ListenAndServe(":3000", nil)
}
```

https://stackoverflow.com/questions/21871479/docker-cant-connect-to-docker-daemon

https://docs.docker.com/engine/install/linux-postinstall/

```bash
sudo groupadd docker
sudo usermod -aG docker $(whoami)
sudo service docker start
docker run hello-world
```

----
```
sudo dapr init
dapr --version
```

--
download golang and install
https://golang.org/doc/install

```
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
```

--
install liteide
get liteidex37.1.linux64-qt5.5.1.tar.gz from github releases page of liteide


--run dapr
```
dapr run --app-id goapp --app-port 3000 --port 3500 ./hello
```

--verify dapr invoke
```
http://localhost:3500/v1.0/invoke/goapp/method/hello
```

**Zero**

run dapr
```
dapr run --app-id zeroapp --app-port 3000 --port 3500 sudo zero
```

wip (does not work)
http://localhost:3500/v1.0/invoke/zeroapp/method/time




