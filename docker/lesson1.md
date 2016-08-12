**TL;DR. Docker in swarm mode can be used as a highly-available and resilient RPC system (no need for Finagle)**

From the same setup in [README](https://github.com/ibmendoza/go-examples/blob/master/docker/README.md), consider the following scenarios:

More about Docker in swarm mode [here](https://t.co/jm5Sxmvl78)

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

```
ID                         NAME          SERVICE     IMAGE       LAST STATE             DESIRED STATE  NODE
aqc05mic310xem0xiylrdbsk3  helloworld.1  helloworld  helloworld  Running 8 minutes ago  Running        manager
6u1u92mdph3c2o9zz222byd4e  helloworld.2  helloworld  helloworld  Running 8 minutes ago  Running        manager
76cg6a69pcb8q1xr42yymhf7d  helloworld.3  helloworld  helloworld  Running 8 minutes ago  Running        manager
esaqpyybzkkxi4s0cpleubcmh  helloworld.4  helloworld  helloworld  Running 8 minutes ago  Running        manager
9wv7h9bmqqfkqyb1kbl335nq8  helloworld.5  helloworld  helloworld  Running 8 minutes ago  Running        manager
```

That is, replicas will be load balanced at the next issuance of ```docker service scale``` command (depending on whether you scale it up or down).

```
root@manager ~# docker service scale helloworld=2
helloworld scaled to 2
root@manager ~# docker service tasks helloworld
ID                         NAME          SERVICE     IMAGE       LAST STATE              DESIRED STATE  NODE
6u1u92mdph3c2o9zz222byd4e  helloworld.2  helloworld  helloworld  Running 10 minutes ago  Running        manager
76cg6a69pcb8q1xr42yymhf7d  helloworld.3  helloworld  helloworld  Running 10 minutes ago  Running        manager
root@manager ~# docker service scale helloworld=5
helloworld scaled to 5
root@manager ~# docker service tasks helloworld
ID                         NAME          SERVICE     IMAGE       LAST STATE               DESIRED STATE  NODE
1bibpse2qxev8sfw3yn81y70j  helloworld.1  helloworld  helloworld  Preparing 2 seconds ago  Running        worker
6u1u92mdph3c2o9zz222byd4e  helloworld.2  helloworld  helloworld  Running 11 minutes ago   Running        manager
76cg6a69pcb8q1xr42yymhf7d  helloworld.3  helloworld  helloworld  Running 11 minutes ago   Running        manager
7ptb7c5zhi6ky1797z0rohmfz  helloworld.4  helloworld  helloworld  Preparing 2 seconds ago  Running        worker
0mwammntx4xse1cg789cpwkv6  helloworld.5  helloworld  helloworld  Preparing 2 seconds ago  Running        worker
```

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

It's the best of both worlds (of course, out of many possible combination).

You can have the speed of bare metal for RPC clients, and the resilience (high availability and load balancing) of containers for RPC backend servers!
