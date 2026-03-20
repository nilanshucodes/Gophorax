//network.go
package main

import (
	"net/http"
	"time"
)

type Result struct {
	Url        string
	StatusCode int
	Latency    float32
}
func Ping(url string) Result {
	start:= time.Now()
	client:= &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)//creates a req object in memory b4 sending it
	//this is why it returns an err immediately—if you pass a malformed URL (like "://google"), it fails here without even touching the network.
	if err != nil {
		return Result{Url: url, StatusCode: 0, Latency: 0}
	}
	req.Header.Set("User-Agent", "GopherPulse/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return Result{Url: url, StatusCode: 0, Latency: 0}
	}
	defer resp.Body.Close()
	latency := float32(time.Since(start).Seconds())
	return Result{Url: url, StatusCode: resp.StatusCode, Latency: latency}
}
//acts as a communication layer
//if web is zombie standard httpget will wait for mins and out worker will be stuck 
