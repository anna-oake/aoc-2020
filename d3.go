package main

import (
	"strconv"
	"strings"
)

func (*methods) D3P1(input string) string {
	lines := strings.Split(input, "\n")

	max := len(lines[0])

	var trees int

	idx := 3
	for i, l := range lines {
		if i == 0 {
			continue
		}
		if string(l[idx]) == "#" {
			trees++
		}
		idx += 3
		if idx >= max {
			idx -= max
		}
	}
	return strconv.Itoa(trees)
}

func (*methods) D3P2(input string) string {
	lines := strings.Split(input, "\n")

	max := len(lines[0])

	var t1, t2, t3, t4, t5 int

	i1 := 1
	i2 := 3
	i3 := 5
	i4 := 7
	i5 := 1
	for n, l := range lines {
		if n == 0 {
			continue
		}

		if string(l[i1]) == "#" {
			t1++
		}
		i1 += 1
		if i1 >= max {
			i1 -= max
		}

		if string(l[i2]) == "#" {
			t2++
		}
		i2 += 3
		if i2 >= max {
			i2 -= max
		}

		if string(l[i3]) == "#" {
			t3++
		}
		i3 += 5
		if i3 >= max {
			i3 -= max
		}

		if string(l[i4]) == "#" {
			t4++
		}
		i4 += 7
		if i4 >= max {
			i4 -= max
		}

		if n%2 == 0 {
			if string(l[i5]) == "#" {
				t5++
			}
			i5 += 1
			if i5 >= max {
				i5 -= max
			}
		}
	}
	return strconv.Itoa(t1 * t2 * t3 * t4 * t5)
}
