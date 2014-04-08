# docker-gocurl
Super-simple Docker container to test `http.Get` from go's `net/http` module because the version check is timing out on some new etcd machines when they try to join a cluster.

### What does it do?
Basically you give it an IP address as an argument and prints the result of doing an `HTTP GET` to `http://IP_ADDRESS_ARG:7001/version`, just like a new etcd machine would try to do when joining a cluster.

### Usage

```
docker run joshuaconner/gocurl IP_ADDRESS_TO_ETCD_MACHINE
```
