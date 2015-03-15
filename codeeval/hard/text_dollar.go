package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convert(n int) string {
	switch {
	case n == 0:
		return ""
	case n <= 9:
		return []string{
			"One", "Two", "Three", "Four", "Five",
			"Six", "Seven", "Eight", "Nine"}[n-1]
	case n <= 19:
		return []string{
			"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
			"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}[n-10]
	case n <= 99:
		return []string{
			"Twenty", "Thirty", "Forty", "Fifty", "Sixty",
			"Seventy", "Eighty", "Ninety"}[n/10-2] + convert(n%10)
	case n <= 999:
		return convert(n/100) + "Hundred" + convert(n%100)
	case n <= 999999:
		return convert(n/1000) + "Thousand" + convert(n%1000)
	case n <= 999999999:
		return convert(n/1000000) + "Million" + convert(n%1000000)
	default:
		return ""
	}
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
		n, _ := strconv.Atoi(string(line))
		fmt.Println(convert(n) + "Dollars")
	}
}
