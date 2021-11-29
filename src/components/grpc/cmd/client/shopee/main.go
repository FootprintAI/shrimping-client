package main

import (
	log "github.com/golang/glog"

	_ "github.com/footprintai/shrimping/components/grpc/version"
)

func main() {
	defer log.Flush()

	Execute()
}
