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

// Create initial array.
func mkArray(x, y int) [][]int32 {
	d := make([][]int32, x)
	for i := range d {
		d[i] = make([]int32, y)
	}
	return d
}

func dumpTable(d [][]int32, x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("% 5d", d[i][j])
		}
		fmt.Println()
	}
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
	p := mkArray(x, y)
	q := mkArray(x, y)

	// Initialize d.
	for i := 0; i < x; i++ {
		if i == 0 {
			d[i][0] = 0
		} else if i == 1 {
			d[i][0] = -8
		} else {
			d[i][0] = -8 - (int32(i) - 1)
		}
	}

	for j := 0; j < y; j++ {
		if j == 0 {
			d[0][j] = 0
		} else if j == 1 {
			d[0][j] = -8
		} else {
			d[0][j] = -8 - (int32(j) - 1)
		}
	}

	// Initialize p.
	for j := 1; j < y; j++ {
		p[0][j] = -999
	}

	// Initialize q.
	for i := 1; i < x; i++ {
		q[i][0] = -999
	}

	// Sequence.
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			p[i][j] = max2(
				d[i-1][j]-8,
				p[i-1][j]-1,
			)
			q[i][j] = max2(
				d[i][j-1]-8,
				q[i][j-1]-1,
			)
			d[i][j] = max3(
				d[i-1][j-1]+score(s1[i-1], s2[j-1]),
				p[i][j],
				q[i][j],
			)
		}
	}
	return d[x-1][y-1]
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
