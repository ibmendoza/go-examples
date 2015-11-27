Note: I just renamed the NSQ clients as lbproducer and lbconsumer.

On VirtualBox, run three Turnkey Linux VMs (using host-only adapter)

```bash
192.168.56.101 - nsqlookupd and lbconsumer (nsq consumer client1)
192.168.56.102 - nsqd and lbproducer (nsq producer client)
192.168.56.103 - lbconsumer (nsq consumer client2)
```

**192.168.56.101**

```
nsqlookupd
```

```
lbconsumer
```

**192.168.56.102**

```bash
--lookupd-tcp-address = IP address of nsqlookupd
--broadcast-address = IP address of machine where nsqd is running
```

```
nsqd --lookupd-tcp-address=192.168.56.101:4160 --broadcast-address=192.168.56.102
```

```
lbproducer
```

**192.168.56.103**

```
lbconsumer --lkp=192.168.56.101
```
