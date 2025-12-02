package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Madjim09/competitiveness.git/internal/works"
)

func main() {
	flag := true
	scanner := bufio.NewScanner(os.Stdin)
	for flag {
		n, err := works.ValidInputIntUser(scanner)
		if err != nil {
			flag = false
			continue
		}
		m, err := works.ValidInputFloatUser(scanner)
		if err != nil {
			flag = false
			continue
		}
		fmt.Println("Ожидание выполнения горутин...")
		data := works.StartGreets(n, m)
		slice := works.ConvChanInSlice(data)
		works.Output(slice)
		flag = works.AnswerUserToQuestion()
	}
}
