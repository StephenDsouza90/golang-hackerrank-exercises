// https://www.hackerrank.com/challenges/python-print/problem
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func printString(number int) {
	var b bytes.Buffer
	for i := 1; i <= number; i++ {
		stringFormat := strconv.Itoa(i)
		b.WriteString(stringFormat)
	}
	fmt.Println(b.String())

}

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error")
	}
	printString(number)
}
