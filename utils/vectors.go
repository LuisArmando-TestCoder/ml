package utils

func GetMultipliedVectors(vector1, vector2 []float64) float64 {
	var multipliedVectors float64
	for i := range vector1 {
		multipliedVectors += vector1[i] * vector2[i]
	}
	return multipliedVectors
}

func GetDotProduct(vector []float64, matrix [][]float64, matrixColumn int) float64 {
	var dotProduct float64
	for vectorIndex, n := range vector {
		dotProduct += n * matrix[vectorIndex][matrixColumn]
	}
	return dotProduct
}

func GetTransposedMatrix(matrix [][]float64) [][]float64 {
	rowSize := len(matrix[0])
	columnSize := len(matrix)
	transposedMatrix := make([][]float64, rowSize)
	for y := range transposedMatrix {
		for i := 0; i < columnSize; i++ {
			transposedMatrix[y] = append(transposedMatrix[y], 0)
		}
	}
	for x, row := range matrix {
		for y, cell := range row {
			transposedMatrix[y][x] = cell
		}
	}
	return transposedMatrix
}

func GetMultipliedMatrix(matrix1, matrix2 [][]float64) [][]float64 {
	transposedMatrix2 := GetTransposedMatrix(matrix2)
	matrixSize := len(matrix1)
	multipliedMatrix := make([][]float64, matrixSize)
	for y := range multipliedMatrix {
		for i := 0; i < matrixSize; i++ {
			multipliedMatrix[y] = append(multipliedMatrix[y], 0)
		}
	}
	for x, row := range matrix1 {
		for y := range transposedMatrix2[x] {
			multipliedMatrix[x][y] = GetDotProduct(row, transposedMatrix2, y)
		}
	}
	return multipliedMatrix
}

func AddFloatingBiasToVector(float float64, vector []float64) []float64 {
	for i := range vector {
		vector[i] += float
	}
	return vector
}