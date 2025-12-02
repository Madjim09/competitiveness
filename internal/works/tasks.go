package works

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type greetVal struct {
	number int
	time   time.Duration
}

func greet(number int, timeSleep float64, sl *[]greetVal, wg *sync.WaitGroup) {
	defer wg.Done()
	var mu sync.Mutex
	start := time.Now()
	time.Sleep(time.Duration(rand.Float64() * timeSleep * float64(time.Second)))
	end := time.Since(start)
	mu.Lock()
	*sl = append(*sl, greetVal{number, end})
	mu.Unlock()
}

func StartGreets(n int, m float64) []greetVal {
	var wg sync.WaitGroup
	wg.Add(n)

	slData := make([]greetVal, 0, n)
	for i := 0; i < n; i++ {
		go greet(i+1, m, &slData, &wg)
	}

	wg.Wait()

	return slData
}

func ConvChanInSlice(data []greetVal) []greetVal {
	sort.Slice(data, func(i, j int) bool {
		if data[i].time == data[j].time {
			return data[i].number < data[j].number
		}
		return data[i].time < data[j].time
	})
	return data
}

func Output(data []greetVal) {
	for _, v := range data {
		fmt.Printf("Горутина номер: %3d выполнилась за %v\n", v.number, v.time)
	}
}

func AnswerUserToQuestion() bool {
	fmt.Print("Хотите продолжить? [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch answer {
		case "y", "yes", "да", "д":
			fmt.Println()
			return true
		case "n", "no", "нет", "н":
			return false
		default:
			fmt.Println("Неверный ввод.")
			fmt.Print("Введите y (да) или n (нет): ")
		}
	}
}
