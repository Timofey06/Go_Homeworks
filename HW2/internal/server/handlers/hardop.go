package handlers

import (
	"math/rand"
	"net/http"
	"time"
)

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	delay := time.Duration(10+rand.Intn(11)) * time.Second
	time.Sleep(delay)

	rand.NewSource(time.Hour.Microseconds())
	code := rand.Intn(13)
	if code != 0 {
		http.Error(w, "Internal Server Error", 499+code)
		return
	}

	w.Write([]byte("OK"))
}
