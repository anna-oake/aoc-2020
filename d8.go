package main

import (
	"strconv"
	"strings"
)

func (*methods) D8P1(input string) string {
	code := strings.Split(input, "\n")

	ran := make(map[int]bool)
	var acc int
	var i int

	for {
		_, present := ran[i]
		if present {
			return strconv.Itoa(acc)
		}
		parts := strings.Split(code[i], " ")
		op := parts[0]
		arg, _ := strconv.Atoi(strings.TrimPrefix(parts[1], "+"))
		ran[i] = true
		switch op {
		case "acc":
			acc += arg
			i += 1
		case "jmp":
			i += arg
		case "nop":
			i += 1
		}
	}
	return "penis"
}

func d8Execute(code []string) (int, bool) {
	ran := make(map[int]bool)
	var acc int
	var i int

	for {
		if i >= len(code) {
			return acc, true
		}
		_, present := ran[i]
		if present {
			return acc, false
		}
		parts := strings.Split(code[i], " ")
		op := parts[0]
		arg, _ := strconv.Atoi(strings.TrimPrefix(parts[1], "+"))
		ran[i] = true
		switch op {
		case "acc":
			acc += arg
			i += 1
		case "jmp":
			i += arg
		case "nop":
			i += 1
		}
	}
}

func (*methods) D8P2(input string) string {
	code := strings.Split(input, "\n")

	for i, cmd := range code {
		parts := strings.Split(cmd, " ")
		op := parts[0]
		if op == "nop" {
			code[i] = "jmp " + parts[1]
			acc, terminated := d8Execute(code)
			if terminated {
				return strconv.Itoa(acc)
			}
			code[i] = cmd
		}
		if op == "jmp" {
			code[i] = "nop " + parts[1]
			acc, terminated := d8Execute(code)
			if terminated {
				return strconv.Itoa(acc)
			}
			code[i] = cmd
		}
	}

	return "penis"
}
