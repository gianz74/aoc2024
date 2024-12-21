package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("cannot read input")
	}

	left, right := splitInput(input)
	score := similarityScore(left, right)
	fmt.Printf("similarity score: %d\n", score)
}

func splitInput(input []byte) ([]int, []int) {
	re := regexp.MustCompile(`([\d]+)[\D]+([\d]+)`)
	matches := re.FindAllSubmatch(input, -1)
	var left, right []int
	for _, line := range matches {
		if len(line) != 3 {
			continue
		}
		numleft, err := strconv.Atoi(string(line[1]))
		if err != nil {
			continue
		}

		left = append(left, numleft)
		numright, err := strconv.Atoi(string(line[2]))
		if err != nil {
			continue
		}

		right = append(right, numright)
	}
	return left, right
}

func similarityScore(left, right []int) int {
	occur := make(map[int]int)
	for _, n := range right {
		occur[n] += 1
	}

	res := 0
	for _, n := range left {
		res += n * occur[n]
	}
	return res
}
