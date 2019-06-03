FROM alpine:3.8

COPY etcd-proxy-server /bin/
CMD ["etcd-proxy-server"]
