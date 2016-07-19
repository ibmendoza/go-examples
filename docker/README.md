Example code: https://github.com/ibmendoza/go-examples/blob/master/docker/helloworld.go

**Run on bare metal**

http://localhost:8080/sdf

Output:

```
Hi there, I love sdf! From: 192.168.0.121 fe80::4c99:7e59:9085:6f87 
fe80::989d:fb29:9977:774c 
fe80::b4ed:4d9b:91e4:d9cd 
192.168.56.1 
192.168.99.1 
192.168.0.121 
```

**With docker run**

First, build the Docker image using the Dockerfile below

```
FROM scratch

COPY helloworld /helloworld

EXPOSE 8080

CMD ["/helloworld"]
```

- upload helloworld Linux binary to Turnkey Linux VM at /home/docker/ipaddr
- chmod +x helloworld
- docker build -t helloworld
- docker images
- docker -p 8080:8080 run helloworld




