package main

import (
	"os"
	"time"

	"github.com/capt4ce/grpc-stream-vs-unary/communication"
	"github.com/capt4ce/grpc-stream-vs-unary/monitoring"
)

func main() {
	var (
		comm       communication.CommunicationInterface
		serverType string
	)
	if len(os.Args) >= 2 {
		serverType = os.Args[1]
	}

	monitoring.StartMonitoring(false)

	if serverType == "stream" {
		serverType = "stream"
	} else {
		serverType = "unary"
	}
	comm = communication.GetCommunicationClient(serverType)

	// sending data to the peer
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < 1000; i++ {
					go func() {
						monitoring.IncrementCounter()
						defer monitoring.DecrementCounter()
						comm.SendRequest()
					}()
				}
			}
		}
	}()

	for {

	}

}
