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
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		// Actual logic
		if prev > 0 {
			if num > prev {
				count++
			}
		}
		prev = num
	}
	fmt.Println(count)
}
