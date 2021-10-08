package main

import (
	"net/http"

	"example.com/ml_tech/routes"
)

func main() {
	http.ListenAndServe(":5000", routes.NewRouter())
}
