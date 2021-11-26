#!/usr/bin/env bash

#npm install -g grpc-tools
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./pb/ --grpc_out=grpc_js:./pb/ *.proto
