package main

import (
	"fmt"
	"strings"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	//
	fmt.Println(strings.Repeat("--", 20))
	PrintMatrixRowWise(matrix)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	PrintMatrixColumnWise(matrix)
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Row Sum of Matrix %v\n", MatrixRowSum(matrix))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Wave Pattern of Row %v\n", WavePrintRow([][]int{{4, 1, 2}, {7, 4, 4}, {3, 7, 4}}))
	fmt.Printf("Wave Pattern of Row %v\n", WavePrintRow([][]int{{1, 2}, {3, 4}}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Wave Pattern of Row %v\n", WavePrintColumn([][]int{{4, 1, 2}, {7, 4, 4}, {3, 7, 4}}))
	fmt.Printf("Wave Pattern of Row %v\n", WavePrintColumn([][]int{{1, 2}, {3, 4}}))
	fmt.Println(strings.Repeat("--", 20))

}
