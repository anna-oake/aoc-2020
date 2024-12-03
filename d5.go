package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (*methods) D5P1(input string) string {
	passes := strings.Split(input, "\n")

	var max int64
	for _, pass := range passes {
		rowS := strings.ReplaceAll(strings.ReplaceAll(pass[:7], "F", "0"), "B", "1")
		row, _ := strconv.ParseInt(rowS, 2, 8)
		seatS := strings.ReplaceAll(strings.ReplaceAll(pass[7:], "L", "0"), "R", "1")
		seat, _ := strconv.ParseInt(seatS, 2, 8)
		id := row*8 + seat
		if id > max {
			max = id
		}
	}
	return strconv.FormatInt(max, 10)
}

func (*methods) D5P2(input string) string {
	passes := strings.Split(input, "\n")

	seats := make(map[int64]string)
	var ids []int64
	for _, pass := range passes {
		rowS := strings.ReplaceAll(strings.ReplaceAll(pass[:7], "F", "0"), "B", "1")
		row, _ := strconv.ParseInt(rowS, 2, 8)
		seatS := strings.ReplaceAll(strings.ReplaceAll(pass[7:], "L", "0"), "R", "1")
		seat, _ := strconv.ParseInt(seatS, 2, 8)
		id := row*8 + seat
		ids = append(ids, id)
		seats[id] = pass
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	for i, id := range ids {
		if i == 0 || i == len(ids)-1 {
			continue
		}
		last := ids[i-1]
		next := ids[i+1]
		if id-1 != last || id+1 != next {
			fmt.Println(id)
		}
	}

	return ""
}
