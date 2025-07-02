package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Data struct {
	Message   string      `json:"message"`
	Count     int         `json:"count"`
	DataAgain []DataAgain `json:"subclass"`
}

type DataAgain struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			data := Data{
				Message: fmt.Sprintf("Item %d", i),
				Count:   i,
				DataAgain: []DataAgain{
					{Message: fmt.Sprintf("Sub %d", j)},
				},
			}
			_ = enc.Encode(data)
			flusher.Flush()
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	http.HandleFunc("/stream", handler)

	log.Println("api jalan ayo test")
	http.ListenAndServe(":8080", nil)

}
