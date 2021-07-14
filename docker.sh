#!/bin/bash

echo 编译镜像...
docker build -t gos/server:v1 .

echo 运行...
docker run --rm -ti -p 8888:8888 \
  -v $PWD/conf:/root/togettoyou/conf \
  -v $PWD/log:/root/togettoyou/log \
  gos/server:v1