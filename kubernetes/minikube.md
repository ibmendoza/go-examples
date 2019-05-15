Installation
============

check if your linux supports virtualization
-------------------------------------------

egrep --color 'vmx|svm' /proc/cpuinfo


download - https://kubernetes.io/docs/tasks/tools/install-minikube/#install-minikube
--------

curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube


add the Minikube executable to your path
----------------------------------------

sudo cp minikube /usr/local/bin && rm minikube


start minikube
--------------

minikube start

start minikube with kvm
-----------------------

minikube start --vm-driver kvm2

https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#kvm2-driver

sudo apt install libvirt-clients libvirt-daemon-system qemu-kvm


https://medium.com/@nieldw/switching-from-minikube-with-virtualbox-to-kvm-2f742db704c9


start minikube without kvm
--------------------------

- https://medium.com/@nieldw/running-minikube-with-vm-driver-none-47de91eab84c
- https://github.com/kubernetes/minikube/blob/master/docs/vmdriver-none.md


sudo minikube start --vm-driver=none --apiserver-ips 127.0.0.1 --apiserver-name localhost
