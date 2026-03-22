// engine.go
package main

import "fmt"

func Worker(id int,jobs<-chan Job,results chan<-Result) {//recieve and send
	for job:= range jobs {//keep receiving work until no more work exists
		fmt.Printf("Worker %d pinging: %s\n", id, job.Url)
		result:= Ping(job.Url)
		results<- result
	}
}
//wihtout for range loop it handle 1 job and exit 
//jobs is a channel reference, we pass the pipe here not water