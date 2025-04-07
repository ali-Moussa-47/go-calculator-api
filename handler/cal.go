package handler

import (
	"calculatorApi/service"
	"fmt"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid numbers", http.StatusBadRequest)
		return
	}
	result, err := service.Calculate(a, b, op)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, `{"result": %d}`, result)
}
