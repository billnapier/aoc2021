package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data")
	if err != nil {
		return
	}

	prev := -1
	a := -1
	b := -1
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		// Actual logic
		if a < 0 {
			a = num
			continue
		}
		if b < 0 {
			b = num
			continue
		}
		sum := num + a + b
		if prev < 0 {
			prev = sum
			continue
		}
		if sum > prev {
			count++
		}
		prev = sum
		a = b
		b = num
	}
	fmt.Println(count)
}
