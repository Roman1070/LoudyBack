package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorWrapper struct {
	Err string `json:"err"`
}

func WriteError(w http.ResponseWriter, err string) {
	errWrapper := ErrorWrapper{Err: err}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.Marshal(errWrapper)
	w.Write(json)
}
