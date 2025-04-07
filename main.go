package main

import (
	"calculatorApi/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calc", handler.CalculateHandler)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
