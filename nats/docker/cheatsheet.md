**Running gnatsd with Docker**

Example from https://medium.com/@itmarketplace.net/clustering-with-nats-f22aeaec7de0

- docker run apcera/gnatsd

- docker network inspect bridge

```
[
    {
        "Name": "bridge",
        "Id": "0e7ebfdd1f27a3971e6e059f2409f9f3a3b907f608eb4b1a70f5cae780f01312",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Internal": false,
        "Containers": {
            "fdf2b6feab62d81f2311535e9085db1586c4862af64daa4f8babba9ed5204169": {
                "Name": "determined_cori",
                "EndpointID": "57cb117537ebcdafa2b7b67e6f90787b3742107e22b13d98f18564a4f177d5e5",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
```

- ./pub -s "nats://172.17.0.2:4222" hello world
