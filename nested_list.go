// https://www.hackerrank.com/challenges/nested-list/problem
package main

import (
	"fmt"
	"sort"
)

type student struct {
	name  string
	grade int
}

var studentGrade = []student{
	{"Harry", 36},
	{"Barry", 60},
	{"Tina", 25},
	{"Akriti", 41},
	{"Harsh", 55},
}

func sortArray(studentGrade []student) {
	// https://yourbasic.org/golang/how-to-sort-in-go/

	sort.SliceStable(studentGrade, func(i, j int) bool {
		return studentGrade[i].grade < studentGrade[j].grade
	})

	fmt.Println(studentGrade[1])

}

func main() {
	sortArray(studentGrade)
}
