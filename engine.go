// engine.go
package main

import "fmt"

func Worker(id int,jobs<-chan Job,results chan<-Result) {
	for job:= range jobs {
		fmt.Printf("Worker %d pinging: %s\n", id, job.Url)
		result:= Ping(job.Url)
		results<- result
	}
}