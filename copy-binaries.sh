#!/usr/bin/env bash

docker create -it --name grpcclient footprintai/shrimping-grpcclient:v1 /bin/bash 
docker cp grpcclient:/out/shrimping.linux shrimping/shrimping.linux
docker cp grpcclient:/out/shrimping.windows.exe shrimping/shrimping.windows.exe
docker cp grpcclient:/out/shrimping.darwin shrimping/shrimping.darwin

docker cp grpcclient:/out/shritagram.linux shritagram/shritagram.linux
docker cp grpcclient:/out/shritagram.windows.exe shritagram/shritagram.windows.exe
docker cp grpcclient:/out/shritagram.darwin shritagram/shritagram.darwin
docker rm -f grpcclient
