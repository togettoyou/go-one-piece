#!/bin/bash

echo 编译镜像...
docker build -t gos/server .

echo 运行...
docker run --rm -v config.yaml:/root/togettoyou/ -p 8888:8888 gos/server