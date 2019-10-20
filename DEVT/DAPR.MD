**Service Discovery**

- multicast DNS - github.com/grandcat/zeroconf (https://github.com/dapr/dapr/blob/master/pkg/discovery/mdns.go)
- makes use of app ID to faciliate service discovery - https://github.com/dapr/docs/tree/master/concepts/publish-subscribe-messaging

**Request/Response (synchronous communication)**

- localhost-oriented programming (a la service mesh)
- within a cluster - https://github.com/dapr/docs/tree/master/howto/invoke-and-discover-services
- using gRPC - https://github.com/dapr/docs/tree/master/howto/create-grpc-app

**Asynchronous Communication**

- [Messaging](https://github.com/dapr/docs/blob/master/concepts/publish-subscribe-messaging/README.md) - follows the Cloud Events spec
- Support Redis Streams, Kafka and cloud-based services ([GCP PubSub](https://github.com/dapr/docs/blob/master/concepts/bindings/README.md), etc) 

**Key/Value Store Support**

- Redis, DynamoDB, CosmosDB, Google Cloud Memorystore (managed Redis service) etc
- Makes use of [state management API](https://github.com/dapr/docs/blob/master/concepts/state-management/state-management.md)
- Dapr supports optimsistic concurrency control (OCC) using ETags.

**Storage**

- S3
- GCP Bucket

**[Distributed Tracing](https://github.com/dapr/docs/blob/master/concepts/distributed-tracing/README.md)**

- based on OpenTelemetry 
- Dapr adds a HTTP/gRPC middleware to the Dapr sidecar. The middleware intercepts all Dapr and application traffic and automatically injects correlation IDs to trace distributed transactions.
- For HTTP requests, Dapr injects a X-Correlation-ID header to requests. For gRPC calls, Dapr inserts a X-Correlation-ID as a field of a header metadata. 
- See reference architecture for distributed tracing at https://linkerd.io/2019/10/07/a-guide-to-distributed-tracing-with-linkerd/

**Secrets Management**

- [Kubernetes secrets](https://github.com/dapr/docs/blob/master/concepts/components/secrets.md)
