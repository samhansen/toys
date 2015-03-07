package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func values(digits string) <-chan int {
	c := make(chan int)
	go func() {
		value, _ := strconv.Atoi(digits)
		c <- value
		value = 0
		for i := 1; i < len(digits); i++ {
			value = value*10 + int(digits[i-1]) - '0'
			for sub := range values(digits[i:]) {
				c <- value + sub
				c <- value - sub
			}
		}
		close(c)
	}()
	return c
}

func is_ugly(value int) bool {
	return (value%2 == 0 ||
		value%3 == 0 ||
		value%5 == 0 ||
		value%7 == 0)
}

func ugly_count(digits string) int {
	count := 0
	for value := range values(digits) {
		if is_ugly(value) {
			count++
		}
	}
	return count
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
		fmt.Println(ugly_count(string(line)))
	}
}
