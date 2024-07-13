package examples

import (
	"math/rand"
)

type Matrix struct {
	matrix  [][]int
	inverse [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(10) // Random integer between 0 and 9
		}
	}

	inverse := make([][]int, cols)
	for i := range inverse {
		inverse[i] = make([]int, rows)
		for j := range inverse[i] {
			inverse[i][j] = matrix[j][i]
		}
	}
	return &Matrix{
		matrix:  matrix,
		inverse: inverse,
	}
}

func (m *Matrix) Execute() error {
	rowsA, colsA := len(m.matrix), len(m.matrix[0])
	_, colsB := len(m.inverse), len(m.inverse[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
		for j := range result[i] {
			for k := 0; k < colsA; k++ {
				result[i][j] += m.matrix[i][k] * m.inverse[k][j]
			}
		}
	}
	return nil
}
