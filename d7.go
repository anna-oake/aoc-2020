package main

import (
	"strconv"
	"strings"
)

func d7CanContain(bags map[string][]string, outer, inner string) bool {
	allowed := bags[outer]
	for _, a := range allowed {
		if a == inner {
			return true
		}
		if d7CanContain(bags, a, inner) {
			return true
		}
	}
	return false
}

func (*methods) D7P1(input string) string {
	rules := strings.Split(input, "\n")

	bags := make(map[string][]string)

	for _, r := range rules {
		parts := strings.Split(r, " bags contain ")
		outer := parts[0]
		if strings.Contains(parts[1], "no other") {
			continue
		}
		parts2 := strings.Split(parts[1], ", ")
		var inner []string
		for _, p := range parts2 {
			p = strings.TrimSuffix(p, ".")
			p = strings.TrimSuffix(p, " bag")
			p = strings.TrimSuffix(p, " bags")
			parts3 := strings.Split(p, " ")
			inner = append(inner, strings.Join(parts3[1:], " "))
		}
		bags[outer] = inner
	}

	var c int
	for outer := range bags {
		if d7CanContain(bags, outer, "shiny gold") {
			c++
		}
	}

	return strconv.Itoa(c)
}

func d7CountInnerBags(bags map[string]map[string]int, outer string) int {
	inner := bags[outer]
	if inner == nil {
		return 0
	}
	var c int
	for bag, count := range inner {
		c += count
		c += count * d7CountInnerBags(bags, bag)
	}
	return c
}

func (*methods) D7P2(input string) string {
	rules := strings.Split(input, "\n")

	bags := make(map[string]map[string]int)

	for _, r := range rules {
		parts := strings.Split(r, " bags contain ")
		outer := parts[0]
		if strings.Contains(parts[1], "no other") {
			continue
		}
		parts2 := strings.Split(parts[1], ", ")
		inner := make(map[string]int)
		for _, p := range parts2 {
			p = strings.TrimSuffix(p, ".")
			p = strings.TrimSuffix(p, " bag")
			p = strings.TrimSuffix(p, " bags")
			parts3 := strings.Split(p, " ")
			count, _ := strconv.Atoi(parts3[0])
			inner[strings.Join(parts3[1:], " ")] = count
		}
		bags[outer] = inner
	}

	return strconv.Itoa(d7CountInnerBags(bags, "shiny gold"))
}
