package helpers_test

import (
	"testing"

	"example.com/ml_tech/helpers"
)

func TestIsMutant(t *testing.T) {
	res := helpers.IsMutant([]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"})
	if res {
		t.Fail()
	}
	t.Log("Success {\"ATzCGA\", \"CAGTGC\", \"TTATTT\", \"AGACGG\", \"GCGTCA\", \"TCACTG\"}")

}
