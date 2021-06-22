// https://www.hackerrank.com/challenges/write-a-function/problem

package main

import (
	"fmt"
)

func isLeap(year int) {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}
		}
	}
}

func main() {
	listOfYears := []int{2000, 2400, 1800, 1900, 2100, 2200, 2300, 2500}
	for i := 0; i <= len(listOfYears)-1; i++ {
		year := listOfYears[i]
		isLeap(year)
	}
}
