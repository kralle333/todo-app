package api

import (
	"fmt"
	"net/http"
)

func (a *todoApp) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	resp := fmt.Sprintf(`{"status": "%s"}`, "all good!")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, resp)
}
