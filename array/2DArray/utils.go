package main

import "fmt"

func PrintMatrixRowWise(list [][]int) {
	row_length := len(list)
	col_length := len((list[0]))
	for i := 0; i < row_length; i++ {
		fmt.Printf("Row %d:", i+1)
		for j := 0; j < col_length; j++ {
			fmt.Printf("%d ", list[i][j])
		}
		fmt.Println()

	}
}

func PrintMatrixColumnWise(list [][]int) {
	row_length := len(list)
	col_length := len((list[0]))
	for i := 0; i < col_length; i++ {
		fmt.Printf("Column %d: ", i)
		for j := 0; j < row_length; j++ {
			fmt.Printf("%d ", list[j][i])
		}
		fmt.Println()

	}
}

func MatrixRowSum(list [][]int) []int {
	row_length := len(list)
	col_length := len((list[0]))

	sum_array := make([]int, row_length)
	for i := 0; i < row_length; i++ {
		row_sum := 0
		for j := 0; j < col_length; j++ {
			row_sum += list[i][j]
		}
		sum_array[i] = row_sum
	}

	return sum_array
}
