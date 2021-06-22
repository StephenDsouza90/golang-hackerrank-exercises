// https://www.hackerrank.com/challenges/python-arithmetic-operators/problem
// https://www.hackerrank.com/challenges/python-division/problem
package main

import "fmt"

func performOperations(a int, b int) {
	add := a + b
	fmt.Println(add)

	difference := a - b
	fmt.Println(difference)

	product := a * b
	fmt.Println(product)

	intDivision := a / b
	fmt.Println(intDivision)

	floatDivision := float64(a) / float64(b)
	fmt.Println(floatDivision)
}

func main() {
	var a int
	var b int
	_, errOne := fmt.Scan(&a)
	checkError(errOne)
	_, errTwo := fmt.Scan(&b)
	checkError(errTwo)

	performOperations(a, b)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}
