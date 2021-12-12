package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data")
	if err != nil {
		return
	}

	var most [12]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '1' {
				most[i] += 1
			} else {
				most[i] -= 1
			}
		}

	}

	// conver to decimal
	x := 0
	for i, s := range most {
		if s > 0 {
			x += 1 << (len(most) - i - 1)
		}
	}

	fmt.Println(x * (x ^ 0b111111111111))
}
