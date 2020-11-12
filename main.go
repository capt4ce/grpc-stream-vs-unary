package main

import (
	"os"

	"github.com/capt4ce/grpc-stream-vs-unary/communication"
	"github.com/capt4ce/grpc-stream-vs-unary/monitoring"
)

func main() {
	var serverType string
	if len(os.Args) >= 2 {
		serverType = os.Args[1]
	}

	monitoring.StartMonitoring(true)

	if serverType == "stream" {
		serverType = "stream"
	} else {
		serverType = "unary"
	}
	communication.StartServer(serverType)

	for {
	}
}
