**Using VirtualBox**

- Install VirtualBox
- Install Turnkey Linux Core
- Set networking to bridged adapter

Commands

1. apt-get update
2. apt-get install apt-transport-https \
     ca-certificates \
     gnupg2 \
     software-properties-common \
     dirmngr --install-recommends \
 3. apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys E04F0923 B3B48BDA
 
 - touch /etc/apt/sources.list
 - apt-get update
 - add-apt-repository "deb https://dl.bintray.com/pony-language/ponylang-debian  $(lsb_release -cs) main"
 
 or manually add below to /etc/apt/sources.list.d/sources.list 
 
 - echo "deb https://dl.bintray.com/pony-language/ponylang-debian stretch main" >> /etc/apt/sources.list.d/sources.list






