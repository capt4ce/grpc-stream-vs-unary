package monitoring

import (
	"runtime"
	"sync"

	"time"

	"github.com/sirupsen/logrus"
)

type Monitor struct {
	count        int
	maxCount     int
	maxGoRoutine int
	sync.Mutex
}

var (
	monitoringInstance Monitor
)

func StartMonitoring(showZeroRequest bool) {
	// ticker := time.NewTicker(50 * time.Millisecond)
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				// count := monitoringInstance.count
				// goRoutineNum := runtime.NumGoroutine()
				// if showZeroRequest || count != 0 {
				// 	logrus.Println("Ongoing requests: ", count)
				// 	logrus.Println("Number of go routines: ", goRoutineNum)
				// 	logrus.Println()
				// }
				logrus.Printf("Highest requests piled up %d", GetMaxCount())
				logrus.Printf("Highest goroutines piled up %d", GetMaxGoroutine())
				logrus.Println()
			}
		}
	}()
}

func IncrementCounter() {
	monitoringInstance.Lock()
	defer monitoringInstance.Unlock()
	monitoringInstance.count++
	if monitoringInstance.count > monitoringInstance.maxCount {
		monitoringInstance.maxCount = monitoringInstance.count
	}
	goRoutineNum := runtime.NumGoroutine()
	if goRoutineNum > monitoringInstance.maxGoRoutine {
		monitoringInstance.maxGoRoutine = goRoutineNum
	}
}

func DecrementCounter() {
	monitoringInstance.Lock()
	defer monitoringInstance.Unlock()
	monitoringInstance.count--
}

func GetMaxCount() int {
	return monitoringInstance.maxCount
}

func GetMaxGoroutine() int {
	return monitoringInstance.maxGoRoutine
}
