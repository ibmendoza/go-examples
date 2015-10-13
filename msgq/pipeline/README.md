### Example in Turnkey Linux

- Create two VMs with VirtualBox Host-only Adapter

#### On first VM (192.168.56.101)

Type:

```sh
./nsqlookupd & ./nsqd --lookupd-tcp-address=127.0.0.1:4160 & ./nsqadmin --lookupd-http-address=127.0.0.1:4161 &
```

#### Assuming pipeline.go has been built with Go

- Still on first VM, open a new Web Shell, then type from respective folder:

```sh
./pipeline node0 tcp://192.168.56.101:40899
```

#### On second VM (192.168.56.102),

```sh
./pipeline node1 tcp://192.168.56.101:40899 "Hello"
```
