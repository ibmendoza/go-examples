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

Computing = input, process, output
Data = input, output

Distributed system = distributed data + distributed computing

### Distributed Data

#### Storage

- Key-value store
- Document-oriented
- Graph
- Relational
- Columnar-oriented

### Distributed Computing

- [FAQ on CAP Theorem](https://henryr.github.io/cap-faq)

- What does asynchronous mean?

An asynchronous network is one in which there is no bound on how long messages may take to be delivered by the network or processed by a machine. The important consequence of this property is that there's no way to distinguish between a machine that has failed, and one whose messages are getting delayed.
