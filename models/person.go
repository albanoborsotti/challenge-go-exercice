package models

type Person struct {
	Id      string   `json:"id"`
	Dna     []string `json:"dna"`
	DnaType string   `json:"dna_type"`
}
