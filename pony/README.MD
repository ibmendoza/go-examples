What is the actor paradigm ?

Reference: https://www.doc.ic.ac.uk/~scd/icooolps15_GC.pdf

The actor paradigm was proposed in 1973 by Carl Hewitt.
An actor is a computational entity that, in response to a message it receives, can: 

- send a finite number of (asynchronous) messages to  other  actors;  
- create  a  finite  number  of  new  actors;  and  
- designate the code to be executed for the next message it receives.

The code executed upon receipt of a message is called a behaviour.

From https://www.reddit.com/r/rust/comments/34rszb/pony_type_and_memory_safe_language/

- graydon

It's very similar to earlier versions of rust. In terms of actor local GC, language provided scheduling of actors, separated function and async invocation system, mixed structural and nominal types, and an overabundance of language supported reference qualifiers. I think they did a good job at it, certainly much better than I did, but I'd be nervous about the cognitive load of the reference qualifiers (capabilities). See this page for example, or the previous few in the tutorial. We had that sort of cognitive load on variables early on and people basically rejected the language because of it, as well as associated compositionality problems. The many years of design iteration had a lot do to with factoring that space into fewer, more essential and general qualifiers. Slow, tedious, world breaking refactoring work. I hope for pony's sake this doesn't happen to them too -- it looks really nice -- but that's my biggest worry for it.

- nikomatsakis

[Caveat: I only skimmed the tutorial and papers.]

I too think it looks like a fun language to use -- but it certainly takes a different tack from Rust. My impression is that, in general, Pony tries to supply up front all the abstractions you will need, whereas Rust is oriented at enabling you to build up safe libraries to fulfill your needs. As Graydon noted, in its early days Rust did take a more Pony-like approach as well, but as ownership/borrowing took shape we found it was more scalable to move that abstraction into libraries.

You can see the contrast in the type systems. Rust has really focused on drilling down to a few core concepts -- ownership and borrows -- and applying those pervasively. This means that we get GC-freedom, data-race-freedom, memory safety, etc, all from the same principles. When we want to extend those ideas, we use library abstractions like Arc<T> to embody new patterns.

Pony, in contrast, has more structure built into the language itself. For data in particular, they have a capabilities system that seems to be the key to their thread-safety guarantees. From what I can tell, data basically begins as "isolated" (ownership) and then gradually decays into other permission modes, eventually settling into immutable or mutable-but-locked-within-one-thread. If you read over the list of capabilities, you can kind of get a picture for the "state machine" they have in mind:

    isolated --> ownership, roughly Box<T>

    value --> immutable, roughly Arc<T>

    reference --> freely aliased within one thread, mutable, roughly Rc<RefCell<T>>

    box --> const (read-only)

    transition --> mutable, but only box aliases

    tag --> opaque pointer

You can see that at various points I highlighted roughly similar Rust data types. This gets at what I meant about Rust preferring to build up handy abstractions in the libraries versus building them into the language. So where in Pony you might take some data you own (e.g., a Foo iso) and transition to the value state (Foo val), in Rust you would create a Foo and then move it into an Arc to get a Arc<Foo>.

Another key difference is that Pony has no notion of borrowing, from what I can see. That is, if you have an isolated pointer in Pony, and you want to make temporary aliases, you have to permanently change to a another state, like value ("immutable"). From that point on, you can alias your structure freely, but you can never mutate it again. In Rust terms, this is the equivalent of moving a Foo into an Arc<Foo>. But I don't see anything equivalent to a temporary reference (&Foo, in Rust terms). (I'll repeat my caveat that this is based on a quick read of the tutorial; perhaps they have a means for temporary borrowing, but I didn't see it in my brief scan -- in any case, I'm sure that they couldn't accommodate "extended borrows" for things like external iterators and so forth.)

In any case, Pony does look like it'd be fun to play with (not that I have time to write much code that's not Rust these days...).
