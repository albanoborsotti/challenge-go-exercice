package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/ml_tech/helpers"
	"example.com/ml_tech/models"
	"example.com/ml_tech/services"
	"github.com/google/uuid"
)

func PostDNA(w http.ResponseWriter, r *http.Request) {
	var body models.DnaRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.Status{
			Status: "Forbidden",
		})
		return
	}

	valid := helpers.ValidateDna(body.Dna)
	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.Status{
			Status: "Forbidden",
		})
		return
	}

	mutant := helpers.IsMutant(body.Dna)
	var person models.Person
	id := uuid.New()
	person.Id = id.String()
	person.Dna = body.Dna
	if mutant {
		person.DnaType = "MUTANT"
	} else {
		person.DnaType = "HUMAN"
	}

	errSave := services.SavePersonDatabase(&person)
	if errSave != nil {
		fmt.Print("Error to save database")
	}
	if mutant {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Status{
			Status: "OK",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(models.Status{
		Status: "Forbidden",
	})
	return

}
