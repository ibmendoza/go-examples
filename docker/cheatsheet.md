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

Unless you specify otherwise with the docker run --net=<NETWORK> option, the Docker daemon connects containers to the bridge network by default.

- https://docs.docker.com/v1.10/engine/userguide/networking/dockernetworks/





