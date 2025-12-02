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
	zero := 0
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
			loc = (start - amt) % 100
		case 'R':
			loc = (start + amt) % 100
		default:
			panic("invalid direction")
		}
		if loc == 0 {
			zero += 1
		}
		start = loc
	}
	fmt.Println(zero)
}
