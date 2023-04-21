FROM golang:1.19 AS builder
WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o etcd-proxy-server ./cmd/etcd-proxy-server/

FROM alpine:latest AS production
WORKDIR /bin
COPY --from=builder /src/etcd-proxy-server .
CMD ["etcd-proxy-server"]