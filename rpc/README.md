**RPC packages** (besides the standard library)

Note: For RPC using Docker Swarm Mode, please see [here](https://github.com/ibmendoza/go-examples/blob/master/docker/lesson1.md).

For data format, prefer to use MessagePack. Just ask [Uber](http://highscalability.com/blog/2016/3/21/to-compress-or-not-to-compress-that-was-ubers-question.html)

For list of data formats, see https://github.com/alecthomas/go_serialization_benchmarks

- https://github.com/tinylib/synapse
- https://github.com/plimble/micro
- https://github.com/valyala/gorpc
- https://github.com/nats-io/nats
- https://github.com/go-mangos/mangos
- https://github.com/ibmendoza/project-iris
- https://github.com/hprose/hprose-go
- https://github.com/grpc/grpc-go
- https://github.com/gorilla/rpc
- https://github.com/zombiezen/go-capnproto2
- https://github.com/smallnest/rpcx
- https://github.com/hashicorp/go-plugin
- https://github.com/funkygao/fae
- https://github.com/ursiform/sleuth
- https://github.com/jondot/armor (based on gRPC)

**RPC Protocols**

- gRPC
- TChannel (https://github.com/uber/tchannel)
- Finagle (https://twitter.github.io/finagle/)
- rpcx (https://github.com/smallnest/rpcx) - based on Go net/rpc

**RPC Proxy**

- Envoy (https://github.com/lyft/envoy) - based on gRPC
- Linkerd (https://github.com/BuoyantIO/linkerd) - based on Finagle

Related links

- https://github.com/hashicorp/net-rpc-msgpackrpc
- https://github.com/cockroachdb/rpc-bench
- https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b
- http://thenewstack.io/grpc-lean-mean-communication-protocol-microservices/
- https://blog.buoyant.io/2016/02/18/linkerd-twitter-style-operability-for-microservices/
- http://www.grpc.io/blog/coreos
- https://github.com/philips/grpc-gateway-example
- http://www.integralist.co.uk/posts/grpc.html
- https://medium.com/square-corner-blog/grpc-cross-platform-open-source-rpc-over-http-2-56c03b5a0173
