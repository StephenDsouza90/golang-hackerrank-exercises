// https://www.hackerrank.com/challenges/python-loops/problem
package main

import (
	"fmt"
	"math"
)

func printSquares(n float64) {
	for i := float64(0); i <= n-1; i++ {
		p := math.Pow(i, 2)
		fmt.Println(p)
	}
}

func main() {
	var n float64
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error")
	}
	printSquares(n)
}
