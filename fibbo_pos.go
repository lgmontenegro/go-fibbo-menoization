package main

import (
	"flag"
)

func main() {
	pos := flag.Int("pos", 0, "Fibbo position required")
	flag.Parse()

	fibboMemoization := make([]int, *pos+1, *pos+1)

	println(fibbo(*pos, fibboMemoization))
}

func fibbo(pos int, fibboMemoization []int) int {
	println("pos: ", pos)
	if fibboMemoization[pos] != 0 {
		println("pos != 0:", fibboMemoization[pos])
		return fibboMemoization[pos]
	}

	if pos <= 2 {
		println("pos <= 2:", 1)
		return 1
	}

	fibboMemoization = append(fibboMemoization, (fibbo(pos-1, fibboMemoization) + fibbo(pos-2, fibboMemoization)))
	return fibboMemoization[pos]
}
