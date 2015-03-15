package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Cartesian coordinate relative to a NxN matrix.
type coordinate struct {
	x, y, n int
}

// Up traverses upwards.  If this would leave the NxN matrix, Up returns the
// origin coordinate.
func (c coordinate) Up() coordinate {
	if c.y == 0 {
		return c
	}
	return coordinate{c.x, c.y - 1, c.n}
}

// Down traverses downwards.  If this would leave the NxN matrix, Down returns
// the origin coordinate.
func (c coordinate) Down() coordinate {
	if c.y == c.n-1 {
		return c
	}
	return coordinate{c.x, c.y + 1, c.n}
}

// Left traverses left.  If this would leave the NxN matrix, Left returns the
// origin coordinate.
func (c coordinate) Left() coordinate {
	if c.x == 0 {
		return c
	}
	return coordinate{c.x - 1, c.y, c.n}
}

// Right traverses right.  If this would leave the NxN matrix, Right returns the
// origin coordinate.
func (c coordinate) Right() coordinate {
	if c.x == c.n-1 {
		return c
	}
	return coordinate{c.x + 1, c.y, c.n}
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max4(a, b, c, d int) int {
	return max2(
		max2(a, b),
		max2(c, d),
	)
}

func score(coord coordinate, s [][]rune, seen string) int {
	// TODO some memoization would probably help.
	char := s[coord.x][coord.y]
	if strings.ContainsRune(seen, char) {
		return 0
	}

	val := 1 + max4(
		score(coord.Up(), s, seen+string(char)),
		score(coord.Down(), s, seen+string(char)),
		score(coord.Left(), s, seen+string(char)),
		score(coord.Right(), s, seen+string(char)),
	)

	return val
}

func makeTable(s string, n int) [][]rune {
	table := make([][]rune, n)
	for i := range table {
		table[i] = make([]rune, n)
	}
	// fill table
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			table[i][j] = rune(s[n*i+j])
		}
	}
	return table
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
		n := int(math.Sqrt(float64(len(line))))
		table := makeTable(string(line), n)
		max := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				max = max2(max, score(coordinate{i, j, n}, table, ""))
			}
		}
		fmt.Println(max)
	}
}
