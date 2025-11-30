package works

import (
	"sync"
	"time"
)

func greet(number int, timeSleep float64, wg *sync.WaitGroup) func() (int, time.Duration) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(time.Duration(timeSleep * float64(time.Second)))
	end := time.Since(start)

	return func() (int, time.Duration) {
		return number, end
	}
}

func StartGreets(n int, m float64) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go greet(i, m, &wg)
	}

	wg.Wait()
}
