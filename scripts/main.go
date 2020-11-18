package main

import (
	"os"
	"sync"
	"time"

	"github.com/capt4ce/grpc-stream-vs-unary/communication"
	"github.com/capt4ce/grpc-stream-vs-unary/monitoring"
	"github.com/sirupsen/logrus"
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

	// // sending requests to the peer with burst requests
	// ticker := time.NewTicker(1 * time.Second)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			for i := 0; i < 100000; i++ {
	// 				go func() {
	// 					monitoring.IncrementCounter()
	// 					defer monitoring.DecrementCounter()
	// 					comm.SendRequest()
	// 				}()
	// 			}
	// 		}
	// 	}
	// }()

	// for {

	// }

	// // sending requests to peer for performance test
	// var wg sync.WaitGroup
	// start := time.Now()
	// for i := 0; i < 100000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		monitoring.IncrementCounter()
	// 		defer monitoring.DecrementCounter()
	// 		defer wg.Done()
	// 		comm.SendRequest()
	// 	}()
	// 	logrus.Println(i)
	// }
	// wg.Wait()
	// elapsed := time.Since(start)
	// logrus.Printf("Requests took %s", elapsed)

	// measuring longest request took
	var wg sync.WaitGroup
	var maxElapsed time.Duration
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			start := time.Now()
			monitoring.IncrementCounter()
			defer func() {
				defer monitoring.DecrementCounter()
				defer wg.Done()
				elapsed := time.Since(start)
				if elapsed > maxElapsed {
					maxElapsed = elapsed
				}
			}()
			comm.SendRequest()
		}()
		logrus.Println(i)
	}
	wg.Wait()
	logrus.Printf("Longest Request(s) took %s", maxElapsed)
}
