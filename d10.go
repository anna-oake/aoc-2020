package main

import (
	"strconv"
	"strings"
)

var d10memo = make(map[int]int64)

func (*methods) D10P1(input string) string {
	lines := strings.Split(input, "\n")
	adapters := make(map[int]bool)
	var target int
	for _, l := range lines {
		j, _ := strconv.Atoi(l)
		adapters[j] = true
		if j > target {
			target = j
		}
	}
	target += 3

	var j, d1, d2, d3 int
	for {
		if target-j <= 3 {
			break
		}
		_, a1 := adapters[j+1]
		if a1 {
			j += 1
			d1++
			continue
		}
		_, a2 := adapters[j+2]
		if a2 {
			j += 2
			d2++
			continue
		}
		_, a3 := adapters[j+3]
		if a3 {
			j += 3
			d3++
			continue
		}
	}
	switch target - j {
	case 3:
		d3++
	case 2:
		d2++
	case 1:
		d1++
	}
	return strconv.Itoa(d1 * d3)
}

func d10connect(adapters map[int]bool, source, target int) int64 {
	if source > target {
		return 0
	}
	if target-source <= 3 {
		return 1
	}
	res, ok := d10memo[source]
	if ok {
		return res
	}
	var cnt, arr int64
	_, a1 := adapters[source+1]
	if a1 {
		arr = d10connect(adapters, source+1, target)
		if arr > 0 {
			cnt += arr
		}
	}
	_, a2 := adapters[source+2]
	if a2 {
		arr = d10connect(adapters, source+2, target)
		if arr > 0 {
			cnt += arr
		}
	}
	_, a3 := adapters[source+3]
	if a3 {
		arr = d10connect(adapters, source+3, target)
		if arr > 0 {
			cnt += arr
		}
	}
	d10memo[source] = cnt
	return cnt
}

func (*methods) D10P2(input string) string {
	lines := strings.Split(input, "\n")
	adapters := make(map[int]bool)
	var target int
	for _, l := range lines {
		j, _ := strconv.Atoi(l)
		adapters[j] = true
		if j > target {
			target = j
		}
	}
	target += 3

	cnt := d10connect(adapters, 0, target)
	return strconv.FormatInt(cnt, 10)
}
