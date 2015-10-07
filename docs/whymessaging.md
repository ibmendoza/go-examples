### Why Messaging?

The following questions are excerpted from Pieter Hintjens' book "ZeroMQ" (only the questions, not the answers)

**How do we handle I/O? Does our application block, or do we handle I/O in the background? This is a key design 
decision. Blocking I/O creates architectures that do not scale well, but background I/O can be very hard to do right**

Using Go sidesteps the issue of blocking I/O with the use of goroutines.

To quote a [StackOverflow](http://stackoverflow.com/questions/6328679/in-golang-does-it-make-sense-to-write-non-blocking-code) post:

> Blocking and non-blocking aren't really about performance, they are about an interface. If you have a single thread 
of execution then a blocking call prevents your program from doing any useful work while it's waiting. But if you have 
multiple threads of execution a blocking call doesn't really matter because you can just leave that thread blocked and 
do useful work in another.

> In Go, a goroutine is swapped out for another one when it blocks on I/O. The Go runtime uses non-blocking I/O syscalls 
to avoid the operating system blocking the thread so a different goroutine can be run on it while the first is waiting 
for it's I/O.

> Goroutines are really cheap so writing non-blocking style code is not needed.
 

**How do we handle dynamic components (i.e., pieces that go away temporarily)? Do we formally split components into 
“clients” and “servers” and mandate that servers cannot disappear? What, then, if we want to connect servers to 
servers? Do we try to reconnect every few seconds?**

Of course, using a messaging broker like Iris or NSQ forces you to split components into clients and servers. With Iris, you are connecting a cluster of servers to another cluster. With NSQ, clients and server need not know each other as long as there is at least one nsqlookupd node running.


**How do we represent a message on the wire? How do we frame data so it’s easy to write and read, safe from buffer 
overflows, and efficient for small messages, yet adequate for the very largest videos of dancing cats wearing party 
hats?**

With Go clients for Iris and NSQ, you can send any message with just a series of bytes. The Go runtime built into the 
compiled Go program manages it for you.


**How do we handle messages that we can’t deliver immediately? Particularly if we’re waiting for a component to come 
back online? Do we discard messages, put them into a database, or put them into a memory queue?**

With Iris, there is no built-in safety net. You have to decide on your own. However, with NSQ, beyond a high-water 
mark, messages are transparently kept on disk.  Messages can be delivered multiple times (client timeouts, 
disconnections, requeues, etc.). It is the client’s responsibility to perform idempotent operations.


**Where do we store message queues? What happens if the component reading from a queue is very slow and causes our 
queues to build up? What’s our strategy then?**

With Iris, messages are stored in memory. If the component reading from a queue is very slow (taking a while to process), there is a possibility of a race condition. That is why, it is recommended that you only use Iris as
message transport and let NSQ do its message queue thing. With NSQ, you can persist messages to disk so you can
re-process it later.


**How do we handle lost messages? Do we wait for fresh data, request a resend, or do we build some kind of reliability layer that ensures messages cannot be lost? What if that layer itself crashes?**

With NSQ, you can persist messages to disk once the high-watermark threshold is reached. That way, you can re-process
those messages later. If NSQ itself crashes, you must architect your client applications to be idempotent. But that
is already a given (I presume).


**What if we need to use a different network transport? Say, multicast instead of TCP unicast? Or IPv6? Do we need to 
rewrite the applications, or is the transport abstracted in some layer?**

With Iris and NSQ, you are limited to TCP transport only. With mangos, there is pluggable support for different
transports.


**How do we route messages? Can we send the same message to multiple peers? Can we send replies back to an original 
requester? How do we write an API for another language? Do we reimplement a wire-level protocol, or do we repackage a 
library? If the former, how can we guarantee efficient and stable stacks? If the latter, how can we guarantee 
interoperability?**

With Iris, you can use four messaging models:

- publish/subscribe
- request/reply
- broadcast
- tunnel

With [mangos](http://bravenewgeek.com/a-look-at-nanomsg-and-scalability-protocols/), you can use six messaging models:

- pair (bidirectional communication)
- request/reply
- pipeline (one-way dataflow)
- bus (many-to-many communication)
- publish/subscribe
- survey (ask a group a question)

With NSQ, you can use publish/subscribe only but then, NSQ has built-in broadcasting and load-balancing of all topics and channels among participating nodes.

With Iris and NSQ, you can write your own client library following its respective protocol besides those already
available.


**How do we represent data so that it can be read between different architectures? Do we enforce a particular encoding for data types? To what extent is this the job of the messaging system rather than a higher layer?**

With Iris, NSQ and mangos, you can represent data simply as bytes. It is up to the higher layer to decode its
representation.


**How do we handle network errors? Do we wait and retry, ignore them silently, or abort?**

With Iris, NSQ and mangos, it is up to you whether message times out or retry. In short, it's an application logic.
