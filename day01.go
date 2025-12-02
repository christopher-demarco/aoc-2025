package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read day01.dat into a slice of strings
func main() {
	dat, err := os.ReadFile("day01.dat")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(dat), "\n")

	start := 50
	wraps := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dir := line[0]
		amt, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		var loc int
		switch dir {
		case 'L':
			// For left rotation, count how many times dial points at 0
			// We visit start-1, start-2, ..., start-amt (mod 100)
			// Count k in [1, amt] where (start - k) ≡ 0 (mod 100), i.e., k ≡ start (mod 100)
			total := start - amt
			loc = total % 100
			if loc < 0 {
				loc += 100
			}
			if start == 0 {
				wraps += amt / 100
			} else if amt >= start {
				wraps += (amt-start)/100 + 1
			}
		case 'R':
			// For right rotation, count wraps going forward
			total := start + amt
			loc = total % 100
			// Count wraps: integer division gives number of times we passed 0
			wraps += total / 100
		default:
			panic("invalid direction")
		}
		start = loc
	}
	fmt.Println(wraps)
}
