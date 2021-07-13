FROM golang:1.14 AS builder
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
COPY . /root/togettoyou/
WORKDIR /root/togettoyou/
RUN make docs

FROM alpine:latest
COPY --from=builder /root/togettoyou/server /root/togettoyou/
COPY --from=builder /root/togettoyou/config.yaml /root/togettoyou/
WORKDIR /root/togettoyou/
EXPOSE 8888
ENTRYPOINT ["./server"]