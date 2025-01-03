package api

import (
	"net/http"
	"strconv"
)

func aboutStatus(w http.ResponseWriter, r *http.Request) {
	method, err := strconv.Atoi(r.URL.Query().Get("method"))

	if err != nil {
		http.Error(w, "Empty method parameter", http.StatusBadRequest)
		return
	}

	res := http.StatusText(method)

	if res == "" {
		http.Error(w, "Unknown method", http.StatusBadRequest)
		return
	}

	w.Write([]byte(res))
}
