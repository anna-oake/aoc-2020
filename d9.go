package main

import (
	"math"
	"strconv"
	"strings"
)

func (*methods) D9P1(input string) string {
	const preamble = 25

	lines := strings.Split(input, "\n")
	var numbers []int64
	for _, l := range lines {
		n, _ := strconv.ParseInt(l, 10, 64)
		numbers = append(numbers, n)
	}
	for i := preamble; i < len(numbers); i++ {
		n := numbers[i]
		m := make(map[int64]bool)
		var valid bool
		for j := i - preamble; j < i; j++ {
			n2 := numbers[j]
			_, ok := m[n2]
			if ok {
				valid = true
				break
			}
			m[n-n2] = true
		}
		if !valid {
			return strconv.FormatInt(n, 10)
		}
	}
	return ""
}

func (*methods) D9P2(input string) string {
	const preamble = 25

	lines := strings.Split(input, "\n")
	var numbers []int64
	for _, l := range lines {
		n, _ := strconv.ParseInt(l, 10, 64)
		numbers = append(numbers, n)
	}
	var invalid int64
	for i := preamble; i < len(numbers); i++ {
		n := numbers[i]
		m := make(map[int64]bool)
		var valid bool
		for j := i - preamble; j < i; j++ {
			n2 := numbers[j]
			_, ok := m[n2]
			if ok {
				valid = true
				break
			}
			m[n-n2] = true
		}
		if !valid {
			invalid = n
			break
		}
	}

	var start, end int
out:
	for i := 0; i < len(numbers)-1; i++ {
		var sum int64
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if sum == invalid {
				start = i
				end = j
				break out
			}
			if sum > invalid {
				break
			}
		}
	}

	var max int64
	min := int64(math.MaxInt64)
	for i := start; i <= end; i++ {
		n := numbers[i]
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return strconv.FormatInt(min+max, 10)
}
