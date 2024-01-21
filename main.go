package matrixcalc

import (
	"errors"	
)


func MultiplyMatrices(matrixA, matrixB [][]int) ([][]int, error) {
	rowsAColumnBA, colsA := len(matrixA), len(matrixA[0])
	rowsAColumnBB, colsB := len(matrixB), len(matrixB[0])

	// Check if matrices can be multiplied
	if colsA != rowsAColumnBB {
		return nil, errors.New("matrices cannot be multiplied")
	}

	// Create the result matrix
	result := make([][]int, rowsAColumnBA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// Perform matrix multiplication
	for i := 0; i < rowsAColumnBA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result, nil
}

