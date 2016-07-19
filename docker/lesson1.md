From the same setup in [README](https://github.com/ibmendoza/go-examples/blob/master/docker/README.md), consider the following scenarios:

**Manager node is up, worker node is down**

- 5 replicas are running

```
http://192.168.0.136:8080/asdf
```

Output:

```
Hi there, I love asdf! From: 10.255.0.9 10.255.0.9 
172.18.0.6 
```

Note: Requests are being load balanced by the Docker engine on a single host (node).


**Manager node is up, worker node is up**

Once the 5 replicas are running on manager node, Docker engine never load balanced the existing replicas to the worker node. However, if you can scale it beyond the existing number of replicas, Docker will load balance some tasks to the worker node.

That is, replicas will be load balanced at the next issuance of ```docker service scale``` command (depending on whether you scale it up or down).

**What does it mean for RPC?**

It means you can build RPC clients that deploys to bare metal and simply reference the IP address of the Docker swarm manager and/or worker nodes. You need not concern about the IP addresses of RPC backend servers.

To illustrate using the same example above,

you can call the RPC server like the following:

```
http://192.168.0.136:8080/asdf
```

or

```
http://192.168.0.137:8080/asdf
```

That is, you only need to know the IP address of the manager node or worker node (assuming the RPC clients are on the same subnet as the manager/worker nodes, aka on the same bridge network as Docker nodes).



