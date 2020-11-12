package main

import (
	"os"

	"github.com/capt4ce/grpc-stream-vs-unary/communication"
)

func main() {
	var serverType string
	if len(os.Args) >= 2 {
		serverType = os.Args[1]
	}

	if serverType == "stream" {
		serverType = "stream"
	} else {
		serverType = "unary"
	}
	communication.StartServer(serverType)

	for {
	}
}
