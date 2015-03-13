package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func score(a, b uint8) int32 {
	if a == b {
		return 3
	}
	return -3
}

// Holds the recursive sublem solutions.  The selected field holds the globally
// optimal solution.
type record struct {
	selected int32 // Globally optimal.
	s1Indel  int32 // Locally optimal assuming s1 ends in an indel.
	s2Indel  int32 // Locally optimal assuming s2 ends in an indel.
}

// Create initial array.
func mkArray(x, y int) [][]record {
	d := make([][]record, x)
	for i := range d {
		d[i] = make([]record, y)
	}
	return d
}

func max2(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int32) int32 {
	if a > max2(b, c) {
		return a
	}
	return max2(b, c)
}

// Dont run on sequnces > 400.
func runTable(s1, s2 string) int32 {
	x := len(s1) + 1
	y := len(s2) + 1

	d := mkArray(x, y)

	// Initialize d.
	for i := 1; i < x; i++ {
		if i == 1 {
			d[i][0].selected = -8
		} else {
			d[i][0].selected = -8 - (int32(i) - 1)
		}
	}

	for j := 1; j < y; j++ {
		if j == 1 {
			d[0][j].selected = -8
		} else {
			d[0][j].selected = -8 - (int32(j) - 1)
		}
	}

	for j := 1; j < y; j++ {
		d[0][j].s1Indel = -999
	}

	for i := 1; i < x; i++ {
		d[i][0].s2Indel = -999
	}

	// Sequence.
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			d[i][j].s1Indel = max2(
				d[i-1][j].selected-8, // Indel start.
				d[i-1][j].s1Indel-1,  // Indel extension.
			)
			d[i][j].s2Indel = max2(
				d[i][j-1].selected-8, // Indel start.
				d[i][j-1].s2Indel-1,  // Indel extension.
			)
			d[i][j].selected = max3(
				d[i-1][j-1].selected+score(s1[i-1], s2[j-1]),
				d[i][j].s1Indel,
				d[i][j].s2Indel,
			)
		}
	}
	return d[x-1][y-1].selected
}

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(fd)

	for {
		line, _, _ := reader.ReadLine()
		if line == nil {
			break
		}
		parts := strings.Split(string(line), "|")
		s1 := strings.TrimSpace(parts[0])
		s2 := strings.TrimSpace(parts[1])
		fmt.Println(runTable(s1, s2))
	}
}
