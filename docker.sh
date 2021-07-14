#!/bin/bash

echo 编译镜像...
docker build -t gos/server:v1 .

echo 运行...
docker run --rm -ti -p 8888:8888 --mount type=bind,source=$PWD/conf,target=/root/togettoyou/conf gos/server