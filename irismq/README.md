IrisMQ = [Iris](https://github.com/ibmendoza/project-iris) and [NSQ](http://nsq.io)

|      | RPC  | Publish/Subscribe                           |Broadcast     |Clustering       |
|:----:|:----:|:-------------------------------------------:|:------------:|:---------------:|
| NSQ  | No   | Broadcast OR Load Balance among consumers   |      No      |   Shard         |
| Iris | Yes  | Broadcast to all consumers                  |      Yes     |P2P (Pastry DHT) |
