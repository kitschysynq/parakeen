package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

type serverConfig struct {
	addr string
}

func main() {
	var cfg serverConfig
	flag.StringVar(&cfg.addr, "addr", ":8080", "Address for server to listen on")
	flag.Parse()

	log.Printf("Listening on %s...", cfg.addr)
	http.HandleFunc("/admin/tests", PostTest)
	log.Fatal(http.ListenAndServe(cfg.addr, nil))
}

func PostTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("bad request from %s, method not allowed: %q", r.RemoteAddr, r.Method)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var t TestRequest
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Printf("error decoding request: %s", err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	log.Printf("ok request from %s", r.RemoteAddr)
	w.WriteHeader(http.StatusNoContent)
}

type TestRequest struct {
	Test struct {
		Result         bool      `json:"result"`
		OrganizationID string    `json:"organization_id"`
		UserID         string    `json:"user_id"`
		Latitude       string    `json:"lat"`
		Longitude      string    `json:"lng"`
		DateTime       time.Time `json:"datetime"`
		Notes          string    `json:"notes"`
	} `json:"test"`
	Image string `json:"test_image"`
}

/*
--data-raw '{
    "test": {
        "result": true,
        "organization_id": "8b2698c1-2cc3-4b0b-94f1-4909872f0dbf",
        "user_id": "4a1a57de-2ab5-4fb5-8850-434d6c55b487",
        "lat": "47.6062",
        "lng": "-122.3321",
        "datetime": "2020-09-21T16:59:09Z",
        "notes": "test upload"
    },
    "test_image": ""
}
*/
