package main

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jsonValue := map[string]any{
			"email":            "oludemiayobami@gmail.con",
			"current_datetime": time.Now().Format("2006-01-02T15:04:05Z"),
			"github_url": "https://github.com/Ayobami0/HNGStage0",
		}
	resp, _ := json.Marshal(jsonValue)

	w.Write(resp)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
