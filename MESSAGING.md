## Messaging-Driven Architecture

- [How PayPal Scaled to Billions of Transactions Daily Using Just 8VMs](http://highscalability.com/blog/2016/8/15/how-paypal-scaled-to-billions-of-transactions-daily-using-ju.html)
- [Why WhatsApp Only Needs 50 Engineers for Its 900M Users](http://www.wired.com/2015/09/whatsapp-serves-900-million-users-50-engineers)
- [Scaling NSQ to 750 Billion Messages](https://segment.com/blog/scaling-nsq)
- [NATS on top of Docker Swarm](http://nats.io/blog/docker-swarm-plus-nats)

## Messaging Patterns

- Request/Reply - (RPC, HTTP on Docker Swarm Mode)
- Publish/Subscribe

**Publish/Subscribe**

- Unicast - point-to-point, one-to-one (e.g. NATS, message transport)
- Groupcast - publish to a group but only one consumer processes the payload (e.g. http://nats.io/documentation/concepts/nats-queueing). See more at https://github.com/IrisMQ/book
- Broadcast - publish to all subscribers / consumers

**Publish/Subscribe Mode**

- pub/sub with no ack (acknowledgement) - tolerate message loss (e.g. NATS)
- pub/sub with ack - at-least-once delivery (e.g. NATS Streaming, NSQ)

## Messaging Patterns (based on [BitMechanic](http://bitmechanic.com/2011/12/30/reasons-to-use-message-queue.html))

- Task Queue (NSQ)
- Delayed Jobs (NSQ consumers)
- Fanout (Publish/subscribe)
- Message Groups (Groupcast)
- RPC ([Docker Swarm mode](https://github.com/ibmendoza/go-examples/blob/master/docker/lesson1.md))

## Message Transport

- https://github.com/project-iris
- https://github.com/go-mangos/mangos
- https://github.com/nats-io
- https://github.com/progrium/go-streamkit

## NSQ and Clients

- https://github.com/nsqio/nsq
- https://github.com/segmentio/go-queue
- https://github.com/tikiatua/nsq-adapter

## PubSub

- https://github.com/chuckpreslar/emission
- https://github.com/olebedev/emitter
- https://github.com/asaskevich/EventBus
- https://github.com/gleicon/swarm

## Others

- https://github.com/adjust/rmq
- https://github.com/dahernan/gopherdiscovery
- https://github.com/joewalnes/websocketd
- https://github.com/centrifugal/centrifugo
- https://github.com/olahol/melody
- https://github.com/KosyanMedia/burlesque
- https://github.com/bogdanovich/siberite
- https://github.com/jcelliott/turnpike
- https://github.com/trevex/golem
- https://github.com/stealthycoin/rhynock
- https://github.com/alecthomas/gorx
- https://github.com/adjust/rmq
- https://github.com/RichardKnop/machinery
- https://github.com/trivago/gollum
- https://github.com/chrislusf/glow
- https://github.com/hashicorp/nomad
- https://github.com/odise/go-cron
- https://github.com/vectaport/flowgraph
- https://github.com/dailymotion/oplog
- https://github.com/jingweno/thunderbird
- https://github.com/LightIO/LightQ

