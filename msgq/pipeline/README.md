### Example in Turnkey Linux

- Create two VMs with VirtualBox Host-only Adapter
- Install NSQ on first VM

#### On first VM (192.168.56.101)

Then type content of runnsq.sh:

#### Assuming pipeline.go has been built with Go

- Still on first VM, open a new Web Shell, then type from respective folder:

```sh
./pipeline node0 tcp://192.168.56.101:40899
```

#### On second VM (192.168.56.102),

```sh
./pipeline node1 tcp://192.168.56.101:40899 "Hello"
```
