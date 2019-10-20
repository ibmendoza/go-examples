**Service Discovery**

- multicast DNS - github.com/grandcat/zeroconf (https://github.com/dapr/dapr/blob/master/pkg/discovery/mdns.go)
- makes use of app ID to faciliate service discovery - https://github.com/dapr/docs/tree/master/concepts/publish-subscribe-messaging

**Request/Response (synchronous communication)**

- localhost-oriented programming (a la service mesh)

**Asynchronous Communication**

- [Messaging](https://github.com/dapr/docs/blob/master/concepts/publish-subscribe-messaging/README.md) - follows the Cloud Events spec
- Support Redis, Kafka and cloud-based services ([GCP PubSub](https://github.com/dapr/docs/blob/master/concepts/bindings/README.md), etc) 

**NoSQL Support**

- Key/Value store - Redis, DynamoDB, CosmosDB, etc

**Storage**

- S3
- GCP Bucket

**[Distributed Tracing](https://github.com/dapr/docs/blob/master/concepts/distributed-tracing/README.md)**

- based on OpenTelemetry 
