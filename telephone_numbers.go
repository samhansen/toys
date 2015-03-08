package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func yield(char rune, s string, c chan<- string) {
	switch char {
	case '1':
		c <- "1" + s
	case '2':
		c <- "a" + s
		c <- "b" + s
		c <- "c" + s
	case '3':
		c <- "d" + s
		c <- "e" + s
		c <- "f" + s
	case '4':
		c <- "g" + s
		c <- "h" + s
		c <- "i" + s
	case '5':
		c <- "j" + s
		c <- "k" + s
		c <- "l" + s
	case '6':
		c <- "m" + s
		c <- "n" + s
		c <- "o" + s
	case '7':
		c <- "p" + s
		c <- "q" + s
		c <- "r" + s
		c <- "s" + s
	case '8':
		c <- "t" + s
		c <- "u" + s
		c <- "v" + s
	case '9':
		c <- "w" + s
		c <- "x" + s
		c <- "y" + s
		c <- "z" + s
	case '0':
		c <- "0" + s
	}
}

func words(digits string) <-chan string {
	c := make(chan string)
	go func() {
		head := rune(digits[0])
		if len(digits) == 1 {
			yield(head, "", c)
		} else {
			for sub := range words(digits[1:]) {
				yield(head, sub, c)
			}
		}
		close(c)
	}()
	return c
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
		all := make([]string, 0, 2187)
		for val := range words(string(line)) {
			all = append(all, val)
		}
		sort.Strings(all)
		fmt.Printf("%s", all[0])
		for i := 1; i < len(all); i++ {
			fmt.Printf(",%s", all[i])
		}
		fmt.Println()
	}
}
