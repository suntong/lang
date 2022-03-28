package main

import (
	_ "expvar"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxWorkers = 5

type job struct {
	name     string
	duration time.Duration
}

func doWork(id int, j job) {
	fmt.Printf("worker%d: started %s, working for %fs\n", id, j.name, j.duration.Seconds())
	time.Sleep(j.duration)
	fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func main() {
	// channel for jobs
	jobs := make(chan job)

	// start workers
	wg := &sync.WaitGroup{}
	wg.Add(maxWorkers)
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			defer wg.Done()

			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	// add jobs
	for i := 0; i < 15; i++ {
		name := fmt.Sprintf("job-%d", i)
		duration := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("adding: %s %s\n", name, duration)
		jobs <- job{name, duration}
	}
	close(jobs)

	// wait for workers to complete
	wg.Wait()
}

/*

$ go run worker_standalone.go
adding: job-0 81ms
adding: job-1 887ms
adding: job-2 847ms
adding: job-3 59ms
adding: job-4 81ms
worker4: started job-3, working for 0.059000s
worker2: started job-0, working for 0.081000s
worker1: started job-1, working for 0.887000s
adding: job-5 318ms
worker5: started job-4, working for 0.081000s
worker3: started job-2, working for 0.847000s
worker4: completed job-3!
worker4: started job-5, working for 0.318000s
adding: job-6 425ms
worker2: completed job-0!
worker2: started job-6, working for 0.425000s
adding: job-7 540ms
worker5: completed job-4!
. . .

*/
