package api

import (
	"math/rand"
	"net/http"
	"strconv"
)

func randomNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit != 0 {
		num := rand.Intn(limit)
		w.Write([]byte(strconv.Itoa(num)))
		return
	}

	num := rand.Intn(10)
	w.Write([]byte(strconv.Itoa(num)))
}
