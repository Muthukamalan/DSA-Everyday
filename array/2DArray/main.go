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

	//
	fmt.Println(strings.Repeat("--", 20))
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	B := [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	if ok, res := AddMatrix(A, B); ok == nil {
		fmt.Printf("Matrix Addtion %v\n", res)
	}
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Larget Element in a Matrix %v\n", LagestInEachRow([][]int{{1, 2}, {1, 3}}))
	fmt.Printf("Larget Element in a Matrix %v\n", LagestInEachRow([][]int{{1, 2, 3}}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("Matrix column wise Sum %v\n", MatrixColSum([][]int{{4, 1}, {1, 3}, {6, 2}}))
	fmt.Printf("Matrix column wise Sum %v\n", MatrixColSum([][]int{{1, 1}, {4, 1}}))
	fmt.Println(strings.Repeat("--", 20))

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("matrix transponse %v\n", MatrixT([][]int{{1, 2, 3}, {4, 5, 7}}))
	fmt.Printf("matrix transponse %v\n", MatrixT([][]int{{4, 1}, {1, 3}, {6, 2}}))
	fmt.Println(strings.Repeat("--", 20))

	//
	A = [][]int{{1, 2, 3}, {4, 5, 6}}
	B = [][]int{{7, 8}, {9, 10}, {11, 12}}
	if ok, res := MatrixMultiply(A, B); ok == nil {
		fmt.Printf("Matrix Mutiplication  %v\n", res)

	}

	C := [][]int{{1, 1, 1}, {1, 1, 4}, {1, 1, 1}}
	D := [][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}
	if ok, res := MatrixMultiply(C, D); ok == nil {
		fmt.Printf("Matrix Mutiplication  %v\n", res)
	}

	//
	fmt.Println(strings.Repeat("--", 20))
	fmt.Printf("is it Identity Matrix %t\n", IsIdentity([][]int{{1, 2, 3}, {4, 5, 7}}))
	fmt.Printf("is it Identity Matrix %t\n", IsIdentity([][]int{{-4, -3}, {-6, 5}}))
	fmt.Printf("is it Identity Matrix %t\n", IsIdentity([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}))
	fmt.Println(strings.Repeat("--", 20))

}
