package main

import (
	"strconv"
	"strings"
)

func (*methods) D6P1(input string) string {
	groups := strings.Split(input, "\n\n")

	var c int
	for _, g := range groups {
		people := strings.Split(g, "\n")
		answers := make(map[string]bool)
		for _, person := range people {
			for _, answer := range person {
				answers[string(answer)] = true
			}
		}
		c += len(answers)
	}

	return strconv.Itoa(c)
}

func (*methods) D6P2(input string) string {
	groups := strings.Split(input, "\n\n")

	var c int
	for _, g := range groups {
		people := strings.Split(g, "\n")
		answers := make(map[string]int)
		for _, person := range people {
			for _, answer := range person {
				answers[string(answer)] = answers[string(answer)] + 1
			}
		}
		for _, q := range answers {
			if q == len(people) {
				c++
			}
		}
	}

	return strconv.Itoa(c)
}
