package main

import (
	"errors"
	"fmt"
)

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

func MatrixColSum(matrix [][]int) []int {
	res := []int{}
	row_length := len(matrix)
	col_length := len((matrix[0]))
	for i := 0; i < col_length; i++ {
		sum_col := 0
		for j := 0; j < row_length; j++ {
			sum_col += matrix[j][i]
		}
		res = append(res, sum_col)
	}

	return res
}

func LagestInEachRow(matrix [][]int) []int {
	row_length := len(matrix)
	col_length := len((matrix[0]))

	array := []int{}
	for i := 0; i < row_length; i++ {
		row_max := matrix[i][0]
		for j := 0; j < col_length; j++ {
			if matrix[i][j] > row_max {
				row_max = matrix[i][j]
			}
		}
		array = append(array, row_max)
	}
	return array
}

func WavePrintRow(matrix [][]int) []int {
	row_length := len(matrix)
	array := []int{}
	for i := 0; i < row_length; i++ {

		col_length := len(matrix[i])

		if i%2 == 0 {
			for j := 0; j < col_length; j++ {
				array = append(array, matrix[i][j])
			}
		} else {
			for j := col_length - 1; j >= 0; j-- {
				array = append(array, matrix[i][j])
			}
		}
	}
	return array

}

func WavePrintColumn(matrix [][]int) []int {
	row_length := len(matrix)
	col_length := len((matrix[0]))

	array := []int{}

	for i := 0; i < col_length; i++ {
		if i%2 == 0 {
			for j := 0; j < row_length; j++ {
				array = append(array, matrix[j][i])
			}
		} else {
			for j := row_length - 1; j >= 0; j-- {
				array = append(array, matrix[j][i])

			}
		}
	}

	return array

}

func AddMatrix(A [][]int, B [][]int) (error, [][]int) {

	matrix_a_row_length := len(A)
	matrix_b_row_length := len(B)

	matrix_a_col_length := len(A[0])
	matrix_b_col_length := len(B[0])
	result := [][]int{}

	if matrix_a_row_length == matrix_b_row_length && matrix_a_col_length == matrix_b_col_length {
		for i := 0; i < matrix_a_row_length; i++ {
			row_sum := []int{}
			for j := 0; j < matrix_b_col_length; j++ {
				row_sum = append(row_sum, A[i][j]+B[i][j])
			}
			result = append(result, row_sum)
		}
		return nil, result
	} else {
		return errors.New("Matrix Size Mismatch"), result

	}

}

func SubMatrix(A [][]int, B [][]int) (error, [][]int) {

	matrix_a_row_length := len(A)

	matrix_a_col_length := len(A[0])
	result := [][]int{}

	if IsMatixSimilarSize(A, B) {
		for i := 0; i < matrix_a_row_length; i++ {
			row_sum := []int{}
			for j := 0; j < matrix_a_col_length; j++ {
				row_sum = append(row_sum, A[i][j]-B[i][j])
			}
			result = append(result, row_sum)
		}
		return nil, result
	} else {
		return errors.New("Matrix Size Mismatch"), result

	}

}

func MatrixMultiply(A, B [][]int) (error, [][]int) {

	rowA := len(A)
	rowB := len(B)
	colA := len(A[0])
	colB := len(B[0])

	if colA != rowB {
		return errors.New("Shape Mismatch"), [][]int{}
	} else {

		// res (rowA x colB )
		res := make([][]int, rowA)
		for i := range res {
			res[i] = make([]int, colB)
		}

		for i := 0; i < rowA; i++ {
			for j := 0; j < colB; j++ {
				for k := 0; k < colA; k++ {
					res[i][j] += A[i][k] * B[k][j]

				}
			}
		}

		return nil, res
	}

}

func MatrixT(matrix [][]int) [][]int {
	matrix_row := len(matrix)
	matrix_col := len(matrix[0])
	res := [][]int{}

	for i := 0; i < matrix_col; i++ {
		each_row := []int{}
		for j := 0; j < matrix_row; j++ {
			each_row = append(each_row, matrix[j][i])
		}
		res = append(res, each_row)
	}

	return res
}

func IsMatixSimilarSize(A, B [][]int) bool {
	matrix_a_row_length := len(A)
	matrix_b_row_length := len(B)

	matrix_a_col_length := len(A[0])
	matrix_b_col_length := len(B[0])

	if matrix_a_row_length == matrix_b_row_length && matrix_a_col_length == matrix_b_col_length {
		return true
	}
	return false
}

func MakeIdentityMatrix(length int) [][]int {
	IdentityMatrix := make([][]int, length)
	for i := 0; i < length; i++ {
		IdentityMatrix[i] = make([]int, length)
		for j := 0; j < length; j++ {
			if i == j {
				IdentityMatrix[i][j] = 1
			} else {
				IdentityMatrix[i][j] = 0
			}
		}
	}
	return IdentityMatrix
}

func IsIdentity(matrix [][]int) bool {
	// matrix (2,4)
	// identity (4,,4)
	// res (2x4)|(4x4) => (2x4)
	IdentityMatrix := MakeIdentityMatrix(len(matrix))
	if ok, res := MatrixMultiply(matrix, IdentityMatrix); ok == nil {
		return isMatrixSiimilar(res, matrix)
	} else {
		return false
	}
}

func isMatrixSiimilar(A, B [][]int) bool {
	if IsMatixSimilarSize(A, B) {
		similar := true
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(A[0]); j++ {
				if A[i][j] != B[i][j] {
					similar = false
				}
			}
		}
		return similar
	}
	return false
}
