package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("cannot read input")
	}

	left, right := splitInput(input)
	slices.Sort(left)
	slices.Sort(right)
	fmt.Printf("distance: %d\n", computeDistance(left, right))
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

func computeDistance(left, right []int) int {
	res := 0
	for i := range left {
		l := left[i]
		r := right[i]
		d := r - l
		if l > r {
			d = -d
		}
		fmt.Printf("i: %d\nleft[%d]: %d, right[%d]: %d\nd: %d\n", i, i, left[i], i, right[i], d)
		res += d
	}
	return res
}
