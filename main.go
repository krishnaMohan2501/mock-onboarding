package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /upi/register", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[MOCK-ONB] registration received | X-Risk-Score=%s", r.Header.Get("X-Risk-Score"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "registration_accepted",
			"userId": "USER-MOCK-001",
		})
	})
	log.Println("[MOCK-ONB] Starting on :9002")
	log.Fatal(http.ListenAndServe(":9002", nil))
}
