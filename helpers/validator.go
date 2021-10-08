package helpers

func ValidateDna(dna []string) bool {
	matrix := BuildMatrix(dna)
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) != len(matrix) {
			return false
		}
		allLeters := contains(matrix[i])
		if !allLeters {
			return false
		}

	}
	return true
}

func contains(s []string) bool {
	count := 0
	for _, v := range s {
		if v == "A" || v == "C" || v == "T" || v == "G" {
			count++
		}
	}
	if count != len(s) {
		return false
	} else {
		return true
	}
}
