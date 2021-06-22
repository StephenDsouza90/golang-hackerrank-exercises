// https://www.hackerrank.com/challenges/py-if-else/problem
package main

import (
	"fmt"
)

func checkN(n int) {
	if n%2 != 0 {
		fmt.Println("Weird - Odd")
	} else if n%2 == 0 {
		rangeOfNumbers(n, 2, 5, "Not Weird - Even")
		rangeOfNumbers(n, 6, 20, "Weird - Even")
		if n > 20 {
			fmt.Println("Not Weird - Even - Greater than 20")
		}
	}
}

func rangeOfNumbers(n int, startValue int, endValue int, message string) {
	for i := startValue; i <= endValue; i++ {
		if n == i {
			fmt.Println(message)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	checkError(err)
	checkN(n)
}
