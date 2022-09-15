package http

import (
	"net/http"
)

func livez(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
