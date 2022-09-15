package healthcheck

import (
	"net/http"
)

func Readyz(w http.ResponseWriter, _ *http.Request) {
	// TODO connect to mysql

	w.WriteHeader(http.StatusOK)
}
