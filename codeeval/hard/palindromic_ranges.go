package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isEven(n int) bool {
	return n&0x01 == 0
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	if length == 1 {
		return true
	}
	if isEven(length) {
		return s[:length/2] == s[length/2:]
	} else {
		midPt := int(math.Floor(float64(length) / 2))
		return s[:midPt] == s[midPt+1:]
	}
}

func ranges(a, b int) int {
	ct := 0
	var palindromes int
	for i := a; i <= b; i++ {
		palindromes = 0
		for j := i; j <= b; j++ {
			if isPalindrome(j) {
				palindromes += 1
			}
			if isEven(palindromes) {
				ct += 1
			}
		}
	}
	return ct
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
		parts := strings.Split(string(line), " ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		fmt.Println(ranges(l, r))
	}
}
