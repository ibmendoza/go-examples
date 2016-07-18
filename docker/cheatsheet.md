**Get IP addr of container**

- http://stackoverflow.com/questions/17157721/getting-a-docker-containers-ip-address-from-the-host

- http://networkstatic.net/10-examples-of-how-to-get-docker-container-ip-address/

- docker inspect container_id

- docker network inspect bridge

**Create overlay network**

- docker network create --driver overlay --subnet=10.0.9.0/24 my-net

- https://docs.docker.com/engine/userguide/networking/get-started-overlay/

**List containers**

- http://stackoverflow.com/questions/16840409/how-do-you-list-containers-in-docker-io

**Understand container networks**

The bridge network represents the ```docker0``` network present in all Docker installations. Unless you specify otherwise with the ```docker run --net=<NETWORK>``` option, the Docker daemon connects containers to this network by default. You can see this bridge as part of a host’s network stack by using the ifconfig command on the host.

Docker does not support automatic service discovery on the default bridge network. If you want to communicate with container names in this default bridge network, you must connect the containers via the legacy ```docker run --link``` option.

However, within a **user-defined bridge network**, linking is not supported. You can expose and publish container ports on containers in this network. 

- https://docs.docker.com/v1.10/engine/userguide/networking/dockernetworks/

- Single-host network - bridge network
- Multi-host network - overlay network






