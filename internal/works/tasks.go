package works

import "time"

func greet(number int, timeSleep float64) func() (int, time.Duration) {
	start := time.Now()
	time.Sleep(time.Duration(timeSleep) * time.Second)
	end := time.Since(start)

	return func() (int, time.Duration) {
		return number, end
	}
}

func StartGreets(n int, m float64) {
	for i := 0; i < n; i++ {
		go greet(i, m)
	}
}
