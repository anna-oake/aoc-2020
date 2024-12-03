package main

import (
	"strconv"
	"strings"
)

func (*methods) D1P1(input string) string {
	lines := strings.Split(input, "\n")
	var num1, num2 int
	for _, line := range lines {
		num1, _ = strconv.Atoi(line)
		for _, otherLine := range lines {
			num2, _ = strconv.Atoi(otherLine)
			if num1+num2 == 2020 {
				return strconv.Itoa(num1 * num2)
			}
		}
	}
	return ""
}

func (*methods) D1P2(input string) string {
	lines := strings.Split(input, "\n")
	var num1, num2, num3 int
	for _, l1 := range lines {
		num1, _ = strconv.Atoi(l1)
		for _, l2 := range lines {
			num2, _ = strconv.Atoi(l2)
			for _, l3 := range lines {
				num3, _ = strconv.Atoi(l3)
				if num1+num2+num3 == 2020 {
					return strconv.Itoa(num1 * num2 * num3)
				}
			}
		}
	}
	return ""
}
