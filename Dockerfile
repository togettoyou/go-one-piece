FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
RUN go env -w GO111MODULE=on

COPY . /root/togettoyou/go-one-server
WORKDIR /root/togettoyou/go-one-server

RUN make docs

EXPOSE 8888
ENTRYPOINT ["./go-one-server"]