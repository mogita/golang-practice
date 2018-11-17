package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomNo int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	// init sum as 0
	sum := 0
	// copy input to local var
	no := number

	// loop while no is not zero
	for no != 0 {
		// set digit to remainder of no mod 10
		// e.g. if no is 8 then digit will be 8
		// e.g. if no is 19 then digit will be 9
		digit := no % 10
		// add digit's value to sum
		sum += digit
		// make no 10 times smaller
		// e.g. if no was 8 then no will be 0.8, but no is int, so actual result shall be 0
		// e.g. if no was 19 then no will be 1.9, but no is int, so actual result shall be 1
		no /= 10
	}

	// wait for 2 seconds
	time.Sleep(2 * time.Second)
	// return the value of sum
	return sum
}

func worker(wg *sync.WaitGroup) {
	// a worker reads from the jobs channel
	for job := range jobs {
		// put current job and its digits() result into a new Result
		output := Result{job, digits(job.randomNo)}
		// write the output to results channel
		results <- output
	}

	// decrease wait group length when all jobs are completed
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	// create a new wait group
	var wg sync.WaitGroup
	// call as many times of worker() as noOfWorkers
	for i := 0; i < noOfWorkers; i++ {
		// increase wait group length
		wg.Add(1)
		// start a worker
		go worker(&wg)
	}

	// wait until all workers are completed
	wg.Wait()
	// close results channel since no more workers are going to write to it
	close(results)
}

func allocate(noOfJobs int) {
	// take noOfJobs to create multiple jobs
	for i := 0; i < noOfJobs; i++ {
		// get a random number at most 998
		randomNo := rand.Intn(999)
		// create a new Job, i as id and randomNo as randomNo
		job := Job{i, randomNo}
		// write the job to jobs channel
		jobs <- job
	}

	// close the jobs channel after writing all jobs
	close(jobs)
}

func result(done chan bool) {
	// a result reads from the results channel
	for result := range results {
		// print the content of the job
		fmt.Printf("Job id %d, input random No %d, sum of digits %d\n", result.job.id, result.job.randomNo, result.sumOfDigits)
	}

	// write to done channel once all results were printed
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken ", diff.Seconds(), " seconds")
}
