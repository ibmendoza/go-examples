### Why Microservices?

#### Your Telco Doesn't Use RPC

First, let me quote a passage from Ross Callon's ["The Twelve Networking Truths"](https://tools.ietf.org/html/rfc1925).

Some things in life can never be fully appreciated nor understood unless experienced firsthand. Some things in 
networking can never be fully understood by someone who neither builds commercial networking equipment nor runs an 
operational network.

If you have never developed a Web or distributed application before, don't despair. Your time will come but when the 
time comes, you better be ready for it.

Second, I want you to familiarize yourself with the [Fallacies of Distributed Computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)
particularly, the first fallacy: "The network is reliable".

Once you have ingrained in your mind that the network is not reliable, eventually things will fall into place.

And third, if you want a thorough exploration of distributed systems, I recommend reading Mikito Takada's ["Distributed
systems"] material at [http://book.mixu.net/distsys/index.html](http://book.mixu.net/distsys/index.html).

Then, once you have a bird's eye view of what we are talking about, it will dawn on you that RPC is certainly [not the 
road you want to go down](https://itjumpstart.wordpress.com/12-rule-app).

So what are microservices?

My definition: microservices = message transport + message queue + Web services.

Regardless whether you are passing messages to a private subnet, on the Internet or through your favorite telecom
network, you are passing it first to a local relay. This is the essence of [overlay network](https://en.wikipedia.org/wiki/Overlay_network).


