### The Minimum You Need to Know about Distributed System

> Distributed programming is the art of solving the same problem that you can solve on a single computer using multiple computers.

> Scalability is the ability of a system, network, or process, to handle a growing amount of work in a capable manner or its 
ability to be enlarged to accommodate that growth

> Performance is characterized by the amount of useful work accomplished by a computer system compared to the time and resources 
used.

> Latency is the state of being latent; delay, a period between the initiation of something and the occurrence.

> Availability is the proportion of time a system is in a functioning condition. If a user cannot access the system, it is said 
to be unavailable. 

> Fault-tolerance is the ability of a system to behave in a well-defined manner once faults occur

via Mikito Takada (http://book.mixu.net/distsys)

> https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing

### Distributed System

- Computing = input, process, output
- Data = input, output

Distributed system = distributed data + distributed computing

### Distributed Data

#### Storage

- Key-value store
- Document-oriented
- Graph
- Columnar-oriented

### Distributed File System

- XtreemFS

### Issues

- Load balancing
- Fault-tolerance (quorum, two-phase commit)
- Cluster membership (gossip, consensus, DHT)
- Consistency vs Availability (consensus vs gossip)
- Data locality vs Remote





### Distributed Computing

- [FAQ on CAP Theorem](https://henryr.github.io/cap-faq)

- What does asynchronous mean?

An asynchronous network is one in which there is no bound on how long messages may take to be delivered by the network or processed by a machine. The important consequence of this property is that there's no way to distinguish between a machine that has failed, and one whose messages are getting delayed.

- What does available mean?

A data store is available if and only if all get and set requests eventually return a response that's part of their specification. This does not permit error responses, since a system could be trivially available by always returning an error.

There is no requirement for a fixed time bound on the response, so the system can take as long as it likes to process a request. But the system must eventually respond.

Notice how this is both a strong and a weak requirement. It's strong because 100% of the requests must return a response (there's no 'degree of availability' here), but weak because the response can take an unbounded (but finite) amount of time.

- What is a partition?

A partition is when the network fails to deliver some messages to one or more nodes by losing them (not by delaying them - eventual delivery is not a partition).

The term is sometimes used to refer to a period during which no messages are delivered between two sets of nodes. This is a more restrictive failure model. We'll call these kinds of partitions total partitions.

The proof of CAP relied on a total partition. In practice, these are arguably the most likely since all messages may flow through one component; if that fails then message loss is usually total between two nodes.
