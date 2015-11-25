
#### Connect nsqd and nsqlookupd in a complete graph

On 192.168.56.101,

./nsqlookupd --tcp-address=192.168.56.101:4160
./nsqd --lookupd-tcp-address=192.168.56.101:4160

On 192.168.56.102,

./nsqlookupd --tcp-address=192.168.56.102:4160
./nsqd --lookupd-tcp-address=192.168.56.101:4160 --lookupd-tcp-address=192.168.56.102:4160

