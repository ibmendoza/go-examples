**RPC packages** (besides the standard library)

Note: As a matter of preference, I like HTTP or TCP-based RPC using Docker Engine in Swarm mode. See it [here](https://github.com/ibmendoza/go-examples/blob/master/docker/lesson1.md).

And using MessagePack as data format! Just ask [Uber](http://highscalability.com/blog/2016/3/21/to-compress-or-not-to-compress-that-was-ubers-question.html)

For data formats, see https://github.com/alecthomas/go_serialization_benchmarks

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

Related links

- https://github.com/hashicorp/net-rpc-msgpackrpc
- https://github.com/cockroachdb/rpc-bench
- https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b
- http://thenewstack.io/grpc-lean-mean-communication-protocol-microservices/
- https://blog.buoyant.io/2016/02/18/linkerd-twitter-style-operability-for-microservices/
- http://www.grpc.io/blog/coreos
- https://github.com/philips/grpc-gateway-example
