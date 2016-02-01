IrisMQ = Iris and NSQ

|      | RPC  | Publish/Subscribe                           |Broadcast     |Clustering       |
|------|-----:|:-------------------------------------------:|:------------:|:---------------:|
| NSQ  | No   | Broadcast OR Load Balance among consumers   |      No      |   Shard         |
| Iris | Yes  | Broadcast to all consumers                  |      Yes     |P2P (Pastry DHT) |
