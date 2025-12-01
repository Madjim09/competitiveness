package works

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ValidInputIntUser(scanner *bufio.Scanner) (int, error) {
	for {
		fmt.Print("Введите количество горутин: ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Ошибка сканера: ", scanner.Err())
			return -1, scanner.Err()
		}
		text := strings.TrimSpace(scanner.Text())
		intVal, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Не верный ввод. Попробуйте еще раз.")
			continue
		}
		return intVal, nil
	}
}

func ValidInputFloatUser(scanner *bufio.Scanner) (float64, error) {
	for {
		fmt.Print("Введите количество секунд (пример: 2; 3.14; 3,14): ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Ошибка сканера: ", scanner.Err())
			return -1, scanner.Err()
		}
		text := strings.TrimSpace(scanner.Text())
		text = strings.Replace(text, ",", ".", 1)
		floatVal, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Не верный ввод. Попробуйте еще раз.")
			continue
		}
		if math.IsNaN(floatVal) || math.IsInf(floatVal, 0) {
			fmt.Println("Введено не число. Попробуйте еще раз.")
			continue
		}
		return floatVal, nil
	}
}
