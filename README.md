# etcd-proxy-server

Proxy to collect [etcd](https://etcd.io/) metrics using [Prometheus](https://prometheus.io/) over HTTPS in a [Kubernetes](https://kubernetes.io/) cluster

![alt text](assets/scheme.png)

* Automatically deployed on master nodes ([daemonSet.yaml](kubernetes/manifests/daemonSet.yaml));
* Using certificates located in `/etc/kubernetes/pki/etcd`, configures the HTTPS client;
* Receives requests for receiving metrics and, on behalf of the configured client, refers to etcd;
* Returns received metrics;
* Allows only `GET` requests for the handler `/metrics`;

## Usage

### Kubernetes

Just do it and forget it:

```sh
git clone git@github.com:eikoshelev/etcd-proxy-server.git && cd etcd-proxy-server
```

```sh
kubectl apply -f kubernetes -n <namespace>
```

### Docker container

```sh
docker run -d --name etcd-proxy-server eikoshelev/etcd-proxy-server:v2.0.0
```

### Description

```
./etcd-proxy-server -h

Usage of ./etcd-proxy-server:
  -caFile string
    	A PEM eoncoded CA's certificate file (default "/etc/kubernetes/pki/etcd/ca.crt")
  -certFile string
    	A PEM eoncoded certificate file (default "/etc/kubernetes/pki/etcd/healthcheck-client.crt")
  -clientTimeout duration
    	Timeout for client (default 10s)
  -hostIP string
    	Host machine IP
  -keyFile string
    	A PEM encoded private key file (default "/etc/kubernetes/pki/etcd/healthcheck-client.key")
  -metricsRoute string
    	Route for collecting metrics (default "/metrics")
  -serverPort string
    	Server port (default ":8888")
  -serverRTimeout duration
    	ReadTimeout for server (default 10s)
  -serverWTimeout duration
    	WriteTimeout for server (default 10s)
```
  
`hostIP` reads the environment variable of the same name by default, which is set depending on the node on which it is deployed under, for more details see [daemonSet.yaml](kubernetes/manifests/daemonSet.yaml#L61)
