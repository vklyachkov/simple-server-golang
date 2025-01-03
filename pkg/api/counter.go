package api

import (
	"net/http"
	"strconv"
)

var Counter int = 0

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Counter++

	w.Write([]byte(strconv.Itoa(Counter)))
}
