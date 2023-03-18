package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func readIDParam(r *http.Request, key string) (int64, error) {
	id := chi.URLParam(r, key)
	return strconv.ParseInt(id, 10, 64)
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

func writeJsonResponse(w http.ResponseWriter, data interface{}) {
	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "marshalling error: %s\n", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshal)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
