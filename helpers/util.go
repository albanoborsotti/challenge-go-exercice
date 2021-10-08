package helpers

import "strings"

func BuildMatrix(dna []string) [][]string {
	var matrix [][]string
	for i := 0; i < len(dna); i++ {
		arr := strings.Split(dna[i], "")
		matrix = append(matrix, arr)
	}
	return matrix
}
