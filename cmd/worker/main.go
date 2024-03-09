package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	concurrency := flag.Int("concurrency", 2, "worker goroutines")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				resp, err := http.Get("http://localhost:7001/dequeue")
				if err != nil || resp.StatusCode == http.StatusNoContent {
					time.Sleep(500 * time.Millisecond)
					continue
				}
				var task map[string]any
				json.NewDecoder(resp.Body).Decode(&task)
				resp.Body.Close()
				log.Printf("worker %d processing %v", id, task["id"])
				fmt.Printf("executed task type=%v\n", task["type"])
			}
		}(i)
	}
	wg.Wait()
}
