- http://jtfmumm.com/blog/2016/03/06/safely-sharing-data-pony-reference-capabilities/

**Laws of shared references**

1) Write Law: Write only when you know no one else can read or write.

2) Read Law: Read only when you know no one else can write.

The Write Law exists because, first, simultaneous writes to the same data structure can lead to unexpected write results and, second, writing while someone else is reading is likely to produce some weird read results. Whether by locks or reference capabilities, the Write Law must be enforced.

The Read Law, similarly, exists because you can only trust the results of a read if you know no one was writing to the data structure at the same time you were trying to read it. And in a perfect world, you could always trust the results of your reads.

**iso reference**

In Pony, if we know that we have the only reference to a mutable data structure, then we have what is called an iso reference, where “iso” stands for “isolated”. An iso is mutable, but it’s safe to send to other actors as long as we give up our reference to it in the process.

**val reference**

But what if we want to be able to continue reading from our map? It’s just that we want to afford other actors the same privilege. In that case, an immutable data structure is called for. Immutable data is the only way to ensure that multiple actors can conform to the Read Law with respect to the same data structure at the same time. In Pony, a reference to an immutable data structure that can be shared among actors is called a val. A val ensures that no one can write to the data structure, though anyone can read from it.
