package main

import "fmt"

func main() {
	fmt.Println(CheckDiv5And7([]int{5, 7, 70, 50, 35}))

}

func CheckDiv5And7(list []int) []int {
	res := []int{}
	for _, i := range list {
		if (i%5 == 0) && (i%7 == 0) {
			res = append(res, i)
		}
	}
	return res
}

// You are given an array A of N elements. You have to make all elements unique. To do so, in one step you can increase any number by one.
// Find the minimum number of steps
func UniqueElement(A []int) int {
	return 0
}
