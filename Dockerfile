FROM golang:1.14 AS builder
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
COPY . /root/togettoyou/go-one-server
WORKDIR /root/togettoyou/go-one-server
RUN make docs

FROM alpine:latest
COPY --from=builder /root/togettoyou/go-one-server/go-one-server /root/togettoyou/go-one-server/
WORKDIR /root/togettoyou/go-one-server
EXPOSE 8888
ENTRYPOINT ["./go-one-server"]