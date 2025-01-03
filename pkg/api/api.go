package api

import (
	"net/http"
)

type api struct {
	addr string
	r    *http.ServeMux
}

// creating a new api instance
func New(addr string, r *http.ServeMux) *api {
	return &api{addr: addr, r: r}
}

// endpoints registration
func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/randomnumber", randomNumber)
	api.r.HandleFunc("/api/count", count)
	api.r.HandleFunc("/api/aboutstatus", aboutStatus)
}

// start server
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}
