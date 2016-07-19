From the same setup in README, consider the following scenarios:

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

