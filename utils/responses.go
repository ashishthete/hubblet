package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
