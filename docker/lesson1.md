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

Note: Requests are being load balanced by the Docker engine.
