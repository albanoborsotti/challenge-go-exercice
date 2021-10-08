package helpers_test

import (
	"testing"

	"example.com/ml_tech/helpers"
)

func TestValidatorOK(t *testing.T) {
	res := helpers.ValidateDna([]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"})
	if !res {
		t.Fail()
	}
	t.Log("Success")
}

func TestValidatorErr(t *testing.T) {
	res := helpers.ValidateDna([]string{"ATzCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"})
	if res {
		t.Fail()
	}
	t.Log("{\"ATzCGA\", \"CAGTGC\", \"TTATTT\", \"AGACGG\", \"GCGTCA\", \"TCACTG\"}", "Is not Valid Schema")
}
