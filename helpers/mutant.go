package helpers

func IsMutant(dna []string) bool {
	matrix := BuildMatrix(dna)
	count := countMatch(matrix)
	if count < 2 {
		return false
	}
	return true
}

func countMatch(matrix [][]string) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		child := matrix[i]
		for c := 0; c < len(child); c++ {
			//Horizontal
			if c < len(child)-3 {
				if child[c] == child[c+1] && child[c] == child[c+2] && child[c] == child[c+3] {
					count++
				}
			}
			//Vertical
			if i < len(matrix)-3 {
				childa := matrix[i+1]
				childb := matrix[i+2]
				childc := matrix[i+3]
				if child[c] == childa[c] && child[c] == childb[c] && child[c] == childc[c] {
					count++
				}
			}
			//Oblicuio
			if i < len(matrix)-3 && c < len(child)-3 {
				childa := matrix[i+1]
				childb := matrix[i+2]
				childc := matrix[i+3]
				if child[c] == childa[c+1] && child[c] == childb[c+2] && child[c] == childc[c+3] {
					count++
				}
			}
		}
	}
	return count
}
