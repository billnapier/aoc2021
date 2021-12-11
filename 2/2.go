package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("data")
	if err != nil {
		return
	}

	depth := 0
	horiz := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		word := parts[0]

		// Actual logic
		if word == "forward" {
			horiz += num
		} else if word == "down" {
			depth += num
		} else if word == "up" {
			depth -= num
		}
	}
	fmt.Println(depth * horiz)
}
