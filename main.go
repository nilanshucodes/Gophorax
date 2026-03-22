// main.go
package main

import (
	"os"
	"fmt"
)
type Job struct {
	Url string
}
//we are using channels to send jobs to workers and get results back
func main(){
	jobs := make(chan Job, 10)//water->Job pipe->jobs
	results := make(chan Result, 10)
	urls:=os.Args[1:]
	numJobs := len(urls)
	go func() {//background task to send jobs to the channel
		for _, url := range urls {
			jobs <- Job{Url: url}
		}
		close(jobs)
	}()
	
	//start workers
	for w:=1; w<=5; w++ {
		go Worker(w, jobs, results)
	}
	var successfulPings int
	var totalLatency float32

	for i:=0; i<numJobs; i++ {
		res := <- results//each time here it pauses ,waits for any worker to send a result back, and then processes it
		if res.StatusCode == 200 {
			successfulPings++
		}
		totalLatency += res.Latency
	}
	fmt.Printf("\n--- GopherPulse Summary ---\n")
	fmt.Printf("Total Checked: %d\n", numJobs)
	fmt.Printf("Successful:    %d\n", successfulPings)
	fmt.Printf("Avg Latency:   %.2fs\n", totalLatency / float32(numJobs))
}
//we are not using basic http.Get this time 
//allow us to set a timeout and a user agent so that sites dont give us a bot title
//main thread can immediately move on to starting workers and collecting results without waiting for each ping to finish