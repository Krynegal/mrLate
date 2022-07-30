package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mrLate/internal/handlers"
	"mrLate/internal/period"
	"net/http"
	"time"
)

func main() {
	client := http.DefaultClient
	timeStr := "2022-08-17T14:50:00+03:00"
	layout := time.RFC3339
	depTime, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(err)
	}
	body := handlers.Request{
		depTime,
		period.Period{
			Hours:   0,
			Minutes: 40,
		},
		period.Period{
			Hours:   0,
			Minutes: 0,
		},
	}
	r, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	log.Println(string(r))
	reqBody := bytes.NewBuffer(r)
	log.Println(reqBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/", reqBody)
	if err != nil {
		panic("bad request")
	}

	resp, err := client.Do(req)
	if err != nil {
		panic("oops... there are some problems with server")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	log.Printf("body: %s", string(respBody))
}
