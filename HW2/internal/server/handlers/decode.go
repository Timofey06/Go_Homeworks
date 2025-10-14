package handlers

import (
	"encoding/base64"
	"encoding/json"
	"hw2/internal/decstr"
	"net/http"
)

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req decstr.DecodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(req.InputString)
	if err != nil {
		http.Error(w, "invalid base64 string", http.StatusBadRequest)
		return
	}

	resp := decstr.DecodeResponse{OutputString: string(decoded)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
