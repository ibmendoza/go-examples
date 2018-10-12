via http://jtfmumm.com/blog/2016/03/06/safely-sharing-data-pony-reference-capabilities/

Safely Sharing Data: Reference Capabilities in Pony
06 Mar 2016

So you’ve got three actors capable of processing data in parallel. You’d like to be able to share data between these actors. What do you do?

If you simply pass around mutable state, then you’re sure to run into situations where your actors start stepping on each other. Actor A is trying to update an entry in a map at the same time Actor B is trying to read that entry. What will Actor B see? Who knows!

So you know what to do: just use locks. Actor B simply has to wait its turn. Nothing could possibly go wrong. But of course locks mean coordination, and coordination is slow. Maybe B didn’t even really care about seeing the latest, greatest version of that entry. Too bad. It’s still going to have to wait.

But we don’t need locks, you say. After all, we have immutability! We’ll just use a persistent hash map. That way B can read its entry no problem and A can simultaneously “write” the update–by creating an altered copy of the immutable map, of course. But if A knows that after sending the map, B will be the only remaining actor who knows about it, too bad! B is still stuck copying on writes. Of course, B could always copy the map into a mutable one, but that doesn’t sound like a recipe for pure, unadulterated speed.

There’s got to be a better way! Well, now there is! And for only 10 monthly payments of $0.00 each, you too can share data faster, safer, and without the headaches. And it’s all because of the magic of Pony reference capabilities.
What is a Reference Capability?

Continuing our earlier example, imagine that we have a hash map with Strings as both keys and values. We’re probably going to store this map as the value of some variable, like so:

  let m = Map[String,String]

What we’ve done here is create a reference called m to the map. This is a mutable map, so we are free to read and write from it as we please.

Zoom out to our multi-actor system, however. If I try to send this data structure to actor B, I encounter a problem. B now has a different reference to the same map. I’m no longer free to read and write as I please. For all I know, B is writing to it at this very moment.

This brings us to two related laws of shared references, laws that you should always keep in mind when thinking about reference capabilities:

1) Write Law: Write only when you know no one else can read or write.

2) Read Law: Read only when you know no one else can write.

The Write Law exists because, first, simultaneous writes to the same data structure can lead to unexpected write results and, second, writing while someone else is reading is likely to produce some weird read results. Whether by locks or reference capabilities, the Write Law must be enforced.

The Read Law, similarly, exists because you can only trust the results of a read if you know no one was writing to the data structure at the same time you were trying to read it. And in a perfect world, you could always trust the results of your reads.

In light of these laws, we see the error in sharing a mutable map with no further thought. Given that the two actors are independent of each other, there’s no way for one to know whether the other is at the moment reading or writing. So without locks, we’re bound for trouble.

What we’d like is a situation where an actor can know for certain whether or not other actors can read from or write to a given data structure. And that’s where reference capabilities come in. A reference capability denies some combination of read and write behaviors, either locally, globally, or both.

The idea should be familiar if you’ve had any experience with immutable data. References to immutable data deny write permissions for everyone. If I’ve shared an immutable data structure, I know I can safely conform to the Read Law since there’s no way anyone else can write to it.

In Pony we get something like this but with a much finer grain of control. Consider, for example, the idea of sharing mutable data. Is this always a bad idea? Not necessarily. In the example above, we created a named reference m to our mutable map. In our local scope, m is the only way we can refer to the map, and hence the only way that we can hope to write to it. What if we wanted to send this mutable map to actor B with no risk of violating our two laws? Do we have to first convert it to an immutable data structure and send that?

In fact we do not. Instead, when sending the map to B, we need only destroy our sole reference to it. If m no longer exists in our local scope, then there is no way we can read from or write to the map. And so it doesn’t concern us in the least whether B is reading from or writing to it.

In Pony, if we know that we have the only reference to a mutable data structure, then we have what is called an iso reference, where “iso” stands for “isolated”. An iso is mutable, but it’s safe to send to other actors as long as we give up our reference to it in the process.

But what if we want to be able to continue reading from our map? It’s just that we want to afford other actors the same privilege. In that case, an immutable data structure is called for. Immutable data is the only way to ensure that multiple actors can conform to the Read Law with respect to the same data structure at the same time. In Pony, a reference to an immutable data structure that can be shared among actors is called a val. A val ensures that no one can write to the data structure, though anyone can read from it.

In an actor system, actors send messages to other actors. In order to send a message to actor C, you have know that C exists. What can C send me to allow me to send to it later? It’s not an iso, for at least two reasons. First, an actor is not the kind of thing you can either read from or write to. Second, C wants many actors to know it exists, so it’s not interested in sending an isolated reference. So what about a val? Well, we’ve already said that you can’t read from an actor, and a val only denies write permissions.

What we need is something that can be shared by many actors but which denies both read and write permissions. In Pony, that’s called a tag. It’s essentially just a reference to the identity of the data in question. If the data is an actor with publicly accessible behaviors, then all we need is a tag to call those behaviors. But we can’t directly read from or write to the object.
The Backbone of Reference Capabilities

As you try to get a handle on how reference capabilities work, it’s helpful to keep in mind the core reason for their existence: safely sharing data across actors. We’re talking about safety we can prove at compile time. And this means that the most important reference capabilities are the sendables. These are references we know we can pass safely to one or more actors: iso, val, and tag, the very three we’ve discussed so far. Think of these as the backbone of Pony reference capabilities.

The sendables stand in the following subtyping relation:

iso –> val –> tag

First, notice that if we expect a tag, we can accept any reference capability. That’s because a tag denies read and write permissions to everyone. So our code that expects a tag will never attempt to read from or write to it, which means there is no chance we’ll violate the Laws of Sharing. So our code can treat both iso and val as subtypes of tag.

Our relation above says that iso is also a subtype of val. Why should this be? Well, whenever we share an iso, we give up our reference to it. And this means that no one but the actor we’re sending to will have a reference. If you’re the only one with a reference, you’re free to decide which reference capability to pick for it. Any code that expects a val can simply treat the iso sent to it as a val. That’s because we can prove at compile time that no one else can write to it, which is exactly what we want for a val.
The Local Toolbox

The sendables are not the only reference capabilities available to us. Pony also provides a toolbox of three reference capabilities that can be useful within the scope of a single actor. These three, called ref, trn, and box, are made possible because of an important property of Pony actors. Within an actor, all execution is serial and synchronous. And this means that within an actor, we do not face the same threats to the Read and Write Laws that we face when concurrent execution is possible. Without concurrency, we can’t have simultaneous writes or a read that occurs during a write.

So in this sense, a single actor need not worry about the Laws of Sharing. The first consequence of this fact is that a single actor is free to hold as many references to a mutable data structure as it needs. A reference to mutable data that makes no guarantees about how many local aliases to that data exist is called a ref. Think of a ref as the old-fashioned reference familiar from programming languages that don’t enforce immutability.

The following is valid Pony code:

  let m: Map[String, String] ref = Map[String, String]
  let n = m
  let o = m

Notice that if m were a Map iso the second line would fail to compile. That’s because we can only have one alias for an iso. For a ref, on the other hand, we can have as many local aliases as we choose.

What if we have a data structure that is currently mutable but which we know will at some future point become immutable. For example, perhaps we are keeping track of votes cast by a fixed number of voters. Once all the votes are in, we will no longer need to mutate our record of them. At that point it will be convenient to be able to freely share the results among actors. In other words, we’ll eventually want a val.

If we used a ref while tallying votes, we’d face a problem. Once all the votes are in, we’d have to find all the existing ref aliases and destroy them. After all, the only way we can prove that it’s safe to share data as a val is if we can show that no one can write to it. Pony provides a cleaner solution here, in the form of the transition reference capability, called trn.

A trn reference is writeable, but allows no other writeable aliases. Unlike an iso, however, it allows other readable aliases. Again, the reason we can allow more than one readable alias to our writeable reference is because we’ve restricted them to the local actor. Hence, we know that we’ll never be reading the data at the same time that we’re writing to it.

The locally readable aliases possess the box reference capability. box means that you can read it locally but not write to it. We may in the course of our vote tallying refer to our box aliases within other data structures (for example, as the field of an object with methods that makes decisions based on the current tally). Once the time comes to “freeze” the vote tally and convert to immutable data, all of these box aliases remain safe to use. We are free to convert our trn to a val.
Converting Between Capabilities

How can we convert a trn to a val? And more generally, how can we convert between capabilities? There are two axes along which these conversions can take place. The first is according to a set of substitution rules that can be modeled in terms of subtyping. The second is by “lifting” reference capabilities, a process called recovery. We’ll begin with substitution.
Substitution

Take a look at the following subtyping diagram:

               --> ref --
              /          \
iso --> trn --            --> box --> tag
              \          /
               --> val --

You can read the arrows in this diagram as “is a subtype of” or “can be substituted for”. For example, iso --> trn can be read as “iso can be substituted for trn”. This subtyping relation is transitive, which means that iso can be substituted for any of the reference capabilities.

Why is this? Recall that if we give up our alias for an iso, then we are free to treat it as any reference capability. How do we give up an alias in Pony? We simply use consume. For example:

  let a: Map[String, String] iso = recover Map[String, String] end
  let b: Map[String, String] trn = consume a

  let c: Map[String, String] iso = recover Map[String, String] end
  let d: Map[String, String] val = consume c

We’ll explain what recover means later, but for now just think of it as a way to show that our Map is an iso.

We can achieve the same effect by consuming our alias when sending an iso. The same reasoning shows us that we can substitute a trn for any reference capability except iso as long as we give up our alias. That’s because trn ensures that there is only one writeable alias in existence and when we consume it, we destroy that alias.

Why can’t we substitute a trn for an iso? Remember that we can have many locally readable box aliases for our trn. This means that when we give up our trn alias, we know that other readable aliases can still exist. And if other aliases exist, then we can’t have an iso.

We can substitute either a ref or a val for a box since both ref and val are locally readable. box doesn’t say that there aren’t writeable aliases anywhere. It just says that the box alias itself is not writeable. So it’s consistent with either locally mutable or globally immutable data. This allows us to write methods that don’t care whether we get a ref or a val as long as the argument is readable.

Finally, for reasons discussed above, any reference capability can be substituted for a tag. That’s because a tag alias is neither readable nor writeable. So no matter what guarantees we want to enforce regarding our subtype, any code that expects a tag will automatically respect them.
Recovery

Subtyping is one way to conceptualize the relationship between reference capabilities. Another is in terms of three distinct categories: mutable, immutable, and opaque.

1) A mutable reference capability denies neither read nor write permissions. This category includes iso, ref, and trn.

2) An immutable reference capability denies write permissions but not read permissions. This category includes val and box.

3) An opaque reference capability denies both read and write permissions. The only example is tag.

It’s worth noting that there is exactly one sendable per category.

Recovery allows us to convert a reference capability to another by doing work within a protected recover block. I call it “protected” because the only aliases external to the block that you can refer to from within the block are sendable ones (i.e. iso, val, and tag). You can create any new aliases you like from within the block (following the normal reference capability rules). The value of a recover block is the value returned at the end of that block. All the other aliases created within the block will be destroyed upon leaving it.

From within a recover block, you can recover the value of the block according to the following three rules:

1) Mutable reference capabilities can be recovered to any mutable, immutable, or opaque reference capabilities. So anything, in other words. For iso and trn, you must consume the alias before you can recover it to something else.

2) Immutable reference capabilities can be recovered to either immutable or opaque reference capabilities.

3) Opaque reference capabilities can be recovered only to opaque. That is, tag can only be recovered to tag, which probably isn’t that useful in practice.

Practically speaking, you can simplify these rules by remembering that immutable reference capabilities can be recovered to immutable or opaque and mutable reference capabilities can be recovered to anything.

What is going on here? Ignoring aliases borrowed from the enclosing scope for the moment, we know that anything that happens within a recover block is invisible to the outside world. So if you create a bunch of ref aliases and return one of those aliases at the end, it’s safe to recover it to an iso. That’s because all the other aliases are destroyed when you leave the block. So we can prove that there is only one reference left. And of course, if there’s only one reference left, then you are free to convert that to any reference capability (as discussed in the substitution section). This explains why a mutable reference created within the block can be recovered to anything.

But you are also allowed to refer to aliases from the enclosing scope, so long as they are sendables. Obviously, if we return a borrowed iso, we’re going to have to consume that alias before we can recover it. That’s because otherwise we’ll have more than one reference to an iso, which is forbidden. So just consume the iso alias when you return it at the end of the block:

  let lst: List[String] iso = recover List[String] end
  let vlst: List[String] val = recover val
    lst.push("hi")
    consume lst
  end

In this example, we start with an iso List. Within the second recover block, we write to it and then return it. But since we consume lst before returning it, we know that there are no aliases left. We are free to recover it to a val as in the example, or anything else for that matter.

So why can’t we recover immutable references to mutable ones? Well, say we borrow a val from the enclosing scope, as in the following (invalid) Pony code:

  let vlst: List[String] val = recover val List[String] end
  ... // we do some stuff with it, which might include sending it
  ... // to someone
  let rlst: List[String] ref = recover ref
    consume vlst
  end

Consuming vlst in the recover block doesn’t do anything useful. That’s because we don’t know how many other aliases exist. After all, we could have sent vlst to other actors. Since it’s a val, we know we can safely share it while holding on to a local reference. So recovering it to a ref violates both of the Laws of Sharing. Everyone else thinks that no one has write permissions, and holding a ref implies that we think no one else has read permissions.

If you contrast those two examples, you’ll notice that when we create an iso, we write:

  recover List[String] end

On the other hand, when we create a val, we write:

  recover val List[String] end

This difference is explained by the default behavior of recover blocks. Remember that there was one sendable per reference capability category: iso for mutable, val for immutable, and tag for opaque. The default behavior of a recover block is to recover a reference capability to the sendable capability in its category. So any mutable reference capability default recovers to iso and any immutable to val.
