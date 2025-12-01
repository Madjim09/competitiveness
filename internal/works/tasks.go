package works

import (
	"bufio"
	"fmt"
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
	sort.Slice(data, func(i, j int) bool {
		return data[i].time < data[j].time
	})
	return data
}

func Output(data []greetVal) {
	for _, v := range data {
		fmt.Printf("Горутина номер: %d выполнилась за: %d\n", v.number, v.time)
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
