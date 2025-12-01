package works

import (
	"sync"
	"time"
)

type greetVal struct {
	number int
	time   time.Duration
}

func greet(number int, timeSleep float64, ch chan greetVal, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(time.Duration(timeSleep * float64(time.Second)))
	end := time.Since(start)
	ch <- greetVal{number, end}
}

func StartGreets(n int, m float64) chan greetVal {
	var wg sync.WaitGroup
	wg.Add(n)

	ch := make(chan greetVal)
	for i := 0; i < n; i++ {
		go greet(i, m, ch, &wg)
	}

	wg.Wait()
	close(ch)
	return ch
}

func ConvChanInSlice(ch chan greetVal) []greetVal {
	data := make([]greetVal, 0, len(ch))
	for v := range ch {
		data = append(data, v)
	}
	return data
}
