package main

import (
	"strconv"
	"strings"
)

func (*methods) D2P1(input string) string {
	lines := strings.Split(input, "\n")

	var valid int

	for _, p := range lines {
		parts := strings.Split(p, ":")
		rules := strings.Split(parts[0], " ")
		minmax := strings.Split(rules[0], "-")
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		letter := rules[1]
		pwd := strings.TrimSpace(parts[1])

		var count int
		for _, c := range pwd {
			if string(c) == letter {
				count++
			}
		}

		if count >= min && count <= max {
			valid++
		}
	}

	return strconv.Itoa(valid)
}

func (*methods) D2P2(input string) string {
	lines := strings.Split(input, "\n")

	var valid int

	for _, p := range lines {
		parts := strings.Split(p, ":")
		rules := strings.Split(parts[0], " ")
		pos := strings.Split(rules[0], "-")
		pos1, _ := strconv.Atoi(pos[0])
		pos2, _ := strconv.Atoi(pos[1])
		letter := rules[1]
		pwd := strings.TrimSpace(parts[1])
		v1 := string(pwd[pos1-1]) == letter
		v2 := string(pwd[pos2-1]) == letter

		if (v1 && !v2) || (!v1 && v2) {
			valid++
		}
	}

	return strconv.Itoa(valid)
}
