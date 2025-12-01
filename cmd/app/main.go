package main

import (
	"bufio"
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
		ch := works.StartGreets(n, m)
		slice := works.ConvChanInSlice(ch)
		works.Output(slice)
		flag = works.AnswerUserToQuestion()
	}
}
