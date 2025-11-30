package works

func greet(c chan int) func() int {
	i := 0
	c <- i
	return func() int {
		i++
		return i
	}
}
