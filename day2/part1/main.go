package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	safe := 0
	for scanner.Scan() {
		row := lineToInts(scanner.Text())
		if isSafe(row) {
			safe++
		}
	}
	fmt.Printf("safe rows: %d\n", safe)
}

func lineToInts(s string) []int {
	re := regexp.MustCompile(`[\d]+`)
	matches := re.FindAll([]byte(s), -1)
	var ret []int

	for _, m := range matches {
		num, err := strconv.Atoi(string(m))
		if err != nil {
			continue
		}
		ret = append(ret, num)
	}
	return ret
}

func isMonotonic(s []int) bool {
	var asc, desc bool
	comp := s[0]
	for _, e := range s {
		if e > comp {
			asc = true
		}
		if e < comp {
			desc = true
		}
		comp = e
	}
	return asc != desc
}

func isGapInRange(s []int) bool {
	var prev int
	for i, curr := range s {
		if i == 0 {
			prev = curr
			continue
		}
		gap := curr - prev
		if gap < 0 {
			gap = -gap
		}
		if gap < 1 {
			return false
		}
		if gap > 3 {
			return false
		}
		prev = curr
	}
	return true
}

func isSafe(s []int) bool {
	return isMonotonic(s) && isGapInRange(s)
}
