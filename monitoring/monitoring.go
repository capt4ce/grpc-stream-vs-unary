package monitoring

import (
	"runtime"
	"sync"

	"time"

	"github.com/sirupsen/logrus"
)

type Monitor struct {
	count int
	sync.Mutex
}

var (
	monitoringInstance Monitor
)

func StartMonitoring() {
	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				count := monitoringInstance.count
				goRoutineNum := runtime.NumGoroutine()
				if count != 0 {
					logrus.Println("Ongoing requests: ", count)
					logrus.Println("Number of go routines: ", goRoutineNum)
					logrus.Println()
				}
			}
		}
	}()
}

func IncrementCounter() {
	monitoringInstance.Lock()
	defer monitoringInstance.Unlock()
	monitoringInstance.count++
}

func DecrementCounter() {
	monitoringInstance.Lock()
	defer monitoringInstance.Unlock()
	monitoringInstance.count--
}
