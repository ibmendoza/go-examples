#### nsqd can be on a node apart from nsqlookupd

192.168.56.101

```bash
nsqlookupd --tcp-address=192.168.56.101:4160
```

192.168.56.102

```bash
nsqd --lookupd-tcp-address=192.168.56.101:4160
```

#### nsqadmin can be on a node apart from nsqlookupd

192.168.56.101

```bash
nsqlookupd --tcp-address=192.168.56.101:4160
```

192.168.56.102

```bash
nsqadmin --lookupd-http-address=192.168.56.101:4161
```

#### nsqd and nsq producer client

nsqd and nsq producer client must be colocated at the same node

#### nsq consumer client and nsqlookupd

nsq consumer client binds to nsqlookupd
