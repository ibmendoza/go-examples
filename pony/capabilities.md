**NOTES:** 

- If you are coming from Go or JavaScript language, you can substitute actor for **goroutine** or **callback** respectively. An **actor** is a unit of async computing.
- The reason for Pony's reference capabilities:  **safely sharing data across actors**
- And this means that the most important reference capabilities are the **sendables**. These are references we know we can pass safely to one or more actors: **1) iso**, **2) val**, and **3) tag**
- Reference: http://jtfmumm.com/blog/2016/03/06/safely-sharing-data-pony-reference-capabilities

**Laws of shared references**

1) Write Law: Write only when you know no one else can read or write.

2) Read Law: Read only when you know no one else can write.

The Write Law exists because, first, simultaneous writes to the same data structure can lead to unexpected write results and, second, writing while someone else is reading is likely to produce some weird read results. Whether by locks or reference capabilities, the Write Law must be enforced.

The Read Law, similarly, exists because you can only trust the results of a read if you know no one was writing to the data structure at the same time you were trying to read it. And in a perfect world, you could always trust the results of your reads.

**iso reference**

In Pony, if we know that we have the only reference to a mutable data structure, then we have what is called an iso reference, where “iso” stands for “isolated”. An iso is mutable, but it’s safe to send to other actors as long as we give up our reference to it in the process.

**val reference**

But what if we want to be able to continue reading from our map? It’s just that we want to afford other actors the same privilege. In that case, an immutable data structure is called for. Immutable data is the only way to ensure that multiple actors can conform to the Read Law with respect to the same data structure at the same time. In Pony, a reference to an immutable data structure that can be shared among actors is called a val. A val ensures that no one can write to the data structure, though anyone can read from it.

**tag reference**

In an actor system, actors send messages to other actors. In order to send a message to actor C, you have know that C exists. What can C send me to allow me to send to it later? It’s not an iso, for at least two reasons. First, an actor is not the kind of thing you can either read from or write to. Second, C wants many actors to know it exists, so it’s not interested in sending an isolated reference. So what about a val? Well, we’ve already said that you can’t read from an actor, and a val only denies write permissions.

What we need is something that can be shared by many actors but which denies both read and write permissions. In Pony, that’s called a tag. It’s essentially just a reference to the identity of the data in question. If the data is an actor with publicly accessible behaviors, then all we need is a tag to call those behaviors. But we can’t directly read from or write to the object.
