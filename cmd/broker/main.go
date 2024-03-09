package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vinir/distributed-work-queue/internal/queue"
)

func main() {
	broker := queue.NewMemoryBroker()
	http.HandleFunc("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Type    string            `json:"type"`
			Payload map[string]string `json:"payload"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		t := broker.Enqueue(req.Type, req.Payload)
		json.NewEncoder(w).Encode(t)
	})
	http.HandleFunc("/dequeue", func(w http.ResponseWriter, r *http.Request) {
		t, ok := broker.Dequeue()
		if !ok {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		json.NewEncoder(w).Encode(t)
	})
	log.Println("broker on :7001")
	log.Fatal(http.ListenAndServe(":7001", nil))
}
