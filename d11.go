package main

import (
	"reflect"
	"strconv"
	"strings"
)

func d11traverse(seats [][]rune, x, y, dir int, nearest bool) rune {
	switch dir {
	case 0:
		y--
	case 1:
		y--
		x++
	case 2:
		x++
	case 3:
		x++
		y++
	case 4:
		y++
	case 5:
		y++
		x--
	case 6:
		x--
	case 7:
		x--
		y--
	}
	if x < 0 || y < 0 || x >= len(seats[0]) || y >= len(seats) {
		return '.'
	}
	s := seats[y][x]
	if nearest {
		return s
	}
	if s == '.' {
		return d11traverse(seats, x, y, dir, nearest)
	}
	return seats[y][x]
}

func d11leave(seats [][]rune, x, y, max int, nearest bool) bool {
	var c int
	for dir := 0; dir < 8; dir++ {
		if d11traverse(seats, x, y, dir, nearest) == '#' {
			c++
		}
		if c == max {
			return true
		}
	}
	return false
}

func d11occupy(seats [][]rune, x, y int, nearest bool) bool {
	for dir := 0; dir < 8; dir++ {
		if d11traverse(seats, x, y, dir, nearest) == '#' {
			return false
		}
	}
	return true
}

func d11(input string, nearest bool, maxOccupied int) string {
	var seats [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row []rune
		for _, seat := range line {
			row = append(row, seat)
		}
		seats = append(seats, row)
	}

	for {
		round := make([][]rune, len(seats))
		for i := range seats {
			round[i] = make([]rune, len(seats[i]))
			copy(round[i], seats[i])
		}

		for y, r := range seats {
			for x, s := range r {
				if s == 'L' {
					if d11occupy(seats, x, y, nearest) {
						round[y][x] = '#'
					}
				}
				if s == '#' {
					if d11leave(seats, x, y, maxOccupied, nearest) {
						round[y][x] = 'L'
					}
				}
			}
		}

		if reflect.DeepEqual(round, seats) {
			var c int
			for _, r := range round {
				for _, s := range r {
					if s == '#' {
						c++
					}
				}
			}
			return strconv.Itoa(c)
		}
		seats = round
	}
	return ""
}

func (*methods) D11P1(input string) string {
	return d11(input, true, 4)
}

func (*methods) D11P2(input string) string {
	return d11(input, false, 5)
}
