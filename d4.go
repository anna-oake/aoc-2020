package main

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func (*methods) D4P1(input string) string {
	var req = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passports := strings.Split(input, "\n\n")

	var count int

	for _, p := range passports {
		valid := true
		for _, r := range req {
			if !strings.Contains(p, r+":") {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return strconv.Itoa(count)
}

func (*methods) D4P2(input string) string {
	var req = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passports := strings.Split(input, "\n\n")

	var count int

	for _, p := range passports {
		valid := true
		for _, r := range req {
			if !strings.Contains(p, r+":") {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		pasp := make(map[string]string)
		p = strings.ReplaceAll(p, "\n", " ")
		fields := strings.Split(p, " ")
		for _, f := range fields {
			kv := strings.Split(f, ":")
			pasp[kv[0]] = kv[1]
		}

		byr, _ := strconv.Atoi(pasp["byr"])
		if byr < 1920 || byr > 2002 {
			continue
		}

		iyr, _ := strconv.Atoi(pasp["iyr"])
		if iyr < 2010 || iyr > 2020 {
			continue
		}

		eyr, _ := strconv.Atoi(pasp["eyr"])
		if eyr < 2020 || eyr > 2030 {
			continue
		}

		if strings.HasSuffix(pasp["hgt"], "cm") {
			hgt, _ := strconv.Atoi(strings.TrimSuffix(pasp["hgt"], "cm"))
			if hgt < 150 || hgt > 193 {
				continue
			}
		} else if strings.HasSuffix(pasp["hgt"], "in") {
			hgt, _ := strconv.Atoi(strings.TrimSuffix(pasp["hgt"], "in"))
			if hgt < 59 || hgt > 76 {
				continue
			}
		} else {
			continue
		}

		if !strings.HasPrefix(pasp["hcl"], "#") {
			continue
		}
		hcl := strings.TrimPrefix(pasp["hcl"], "#")
		if len(hcl) != 6 {
			continue
		}
		_, err := hex.DecodeString(hcl)
		if err != nil {
			continue
		}

		eyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		valid = false
		for _, e := range eyes {
			if pasp["ecl"] == e {
				valid = true
			}
		}
		if !valid {
			continue
		}

		if len(pasp["pid"]) != 9 {
			continue
		}
		pid, _ := strconv.Atoi(pasp["pid"])
		if pid < 1 {
			continue
		}

		count++
	}
	return strconv.Itoa(count)
}
