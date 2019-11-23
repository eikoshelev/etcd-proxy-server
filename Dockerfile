FROM golang:alpine AS build
WORKDIR /src
ADD . .
WORKDIR /src
RUN go build -o etcd-proxy-server

FROM alpine
WORKDIR /bin
COPY --from=build /src/etcd-proxy-server .
CMD ["etcd-proxy-server"]
