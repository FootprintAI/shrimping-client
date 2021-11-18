#!/usr/bin/env bash

docker create -it --name grpcclient footprintai/shrimping-grpcclient:v1 /bin/bash 
docker cp grpcclient:/out/shrimping.linux binary/shrimping.linux
docker cp grpcclient:/out/shrimping.windows.exe binary/shrimping.windows.exe
docker cp grpcclient:/out/shrimping.darwin binary/shrimping.darwin
docker rm -f grpcclient
