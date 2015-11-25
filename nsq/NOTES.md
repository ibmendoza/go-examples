#### nsqadmin can be on a node apart from nsqlookupd

192.168.56.101

```bash
nsqlookupd --tcp-address=192.168.56.101:4160
```

192.168.56.102

```bash
nsqadmin --lookupd-http-address=192.168.56.101:4161
```
