https://raesene.github.io/blog/2016/07/23/Docker-MacVLAN/

https://github.com/docker/libnetwork/blob/master/docs/macvlan.md

```bash
$ ifconfig
```
```
docker0   Link encap:Ethernet  HWaddr 02:42:08:6d:21:65  
          inet addr:172.17.0.1  Bcast:0.0.0.0  Mask:255.255.0.0
          UP BROADCAST MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

docker_gwbridge Link encap:Ethernet  HWaddr 02:42:6d:4b:05:03  
          inet addr:172.18.0.1  Bcast:0.0.0.0  Mask:255.255.0.0
          inet6 addr: fe80::42:6dff:fe4b:503/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:8 errors:0 dropped:0 overruns:0 frame:0
          TX packets:39 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:536 (536.0 B)  TX bytes:4967 (4.9 KB)

enp0s25   Link encap:Ethernet  HWaddr 00:21:85:a2:3f:6f  
          inet addr:192.168.254.102  Bcast:192.168.254.255  Mask:255.255.255.0
          inet6 addr: fe80::ea96:b753:e556:7e79/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:6709 errors:0 dropped:0 overruns:0 frame:0
          TX packets:5874 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:6716953 (6.7 MB)  TX bytes:627821 (627.8 KB)
          Interrupt:20 Memory:fe7c0000-fe7e0000 

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:1517 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1517 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1 
          RX bytes:191689 (191.6 KB)  TX bytes:191689 (191.6 KB)

vetha866bf8 Link encap:Ethernet  HWaddr 4a:6b:30:9b:ee:aa  
          inet6 addr: fe80::486b:30ff:fe9b:eeaa/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:8 errors:0 dropped:0 overruns:0 frame:0
          TX packets:65 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:648 (648.0 B)  TX bytes:8163 (8.1 KB)
```

```bash
sudo docker network create -d macvlan 
--subnet=192.168.254.0/24 
--gateway=192.168.254.1 
--ip-range=192.168.254.128/26 
-o parent=enp0s25 testnet
```

9f558575c1d9b8b3fa4a9d36ca873c3b802adf4aac217721591b2d1b9fa84e09

```bash
$ docker version
```
```sh
Client:
 Version:      1.12.4
 API version:  1.24
 Go version:   go1.6.4
 Git commit:   1564f02
 Built:        Tue Dec 13 00:08:34 2016
 OS/Arch:      linux/amd64
```

```bash
$ service docker status
```

```sh
● docker.service - Docker Application Container Engine
   Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset:
   Active: active (running) since Wed 2016-12-21 14:29:58 PHT; 1min 38s ago
     Docs: https://docs.docker.com
 Main PID: 2521 (dockerd)
    Tasks: 19
   Memory: 57.8M
      CPU: 974ms
   CGroup: /system.slice/docker.service
           ├─2521 /usr/bin/dockerd -H fd://
           └─3162 docker-containerd -l unix:///var/run/docker/libcontainerd/do
```

```bash
$ sudo docker run -it --network=testnet alpine /bin/sh
```

```sh
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
3690ec4760f9: Pull complete 
Digest: sha256:1354db23ff5478120c980eca1611a51c9f2b88b61f24283ee8200bf9a54f2e5c
Status: Downloaded newer image for alpine:latest
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:C0:A8:FE:80  
          inet addr:192.168.254.128  Bcast:0.0.0.0  Mask:255.255.255.0
          inet6 addr: fe80::42:c0ff:fea8:fe80%32702/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:6 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:86 (86.0 B)  TX bytes:508 (508.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1%32702/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
/ #                                                                                                   
```
