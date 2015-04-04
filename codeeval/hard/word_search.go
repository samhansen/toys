package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	X, Y int
}

var board = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'},
}

func contains(s []byte, c coord, seen map[coord]bool) bool {
	if len(s) == 0 {
		return true
	}
	_, ok := seen[c]
	if ok {
		return false
	}
	prefix := s[0]
	suffix := s[1:]
	if prefix != board[c.X][c.Y] {
		return false
	}
	seen[c] = true
	var left, right, up, down bool
	if c.X > 0 {
		up = contains(suffix, coord{c.X - 1, c.Y}, seen)
	}
	if c.X < 2 {
		down = contains(suffix, coord{c.X + 1, c.Y}, seen)
	}
	if c.Y > 0 {
		left = contains(suffix, coord{c.X, c.Y - 1}, seen)
	}
	if c.Y < 3 {
		right = contains(suffix, coord{c.X, c.Y + 1}, seen)
	}
	return left || right || up || down
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
		var valid bool
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				seen := make(map[coord]bool)
				if contains(line, coord{i, j}, seen) {
					valid = true
				}
			}
		}
		if valid {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
