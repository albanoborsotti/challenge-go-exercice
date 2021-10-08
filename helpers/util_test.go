package helpers_test

import (
	"testing"

	"example.com/ml_tech/helpers"
)

func TestUtils(t *testing.T) {
	matrix := helpers.BuildMatrix([]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"})
	var exp [][]string
	exp = append(exp, []string{"A", "T", "G", "C", "G", "A"})
	exp = append(exp, []string{"C", "A", "G", "T", "G", "C"})
	exp = append(exp, []string{"T", "T", "A", "T", "T", "T"})
	exp = append(exp, []string{"A", "G", "A", "C", "G", "G"})
	exp = append(exp, []string{"G", "C", "G", "T", "C", "A"})
	exp = append(exp, []string{"T", "C", "A", "C", "T", "G"})
	if matrix[0][0] != exp[0][0] {
		t.Fail()
	}
	t.Log("Sucees")

}
