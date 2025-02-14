package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ErrorWrapper struct {
	Err string `json:"err"`
}

const DefaultSessionLifetime = 10 * time.Hour

var DefaultSessionLifetimeString = strconv.Itoa(int(DefaultSessionLifetime.Seconds()))

func WriteError(w http.ResponseWriter, err string) {
	errWrapper := ErrorWrapper{Err: err}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.Marshal(errWrapper)
	w.Write(json)
}
