// See
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://gist.github.com/harlow/dbcd639cf8d396a2ab73

package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type job struct {
	name     string
	duration time.Duration
}

func doWork(id int, j job) {
	fmt.Printf("worker%d: started %s, working for %f seconds\n", id, j.name, j.duration.Seconds())
	time.Sleep(j.duration)
	fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func requestHandler(jobs chan job, w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the durations.
	duration, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate delay is in range 1 to 10 seconds.
	if duration.Seconds() < 1 || duration.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// Set name and validate value.
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	// Create Job and push the work onto the jobCh.
	job := job{name, duration}
	go func() {
		fmt.Printf("added: %s %s\n", job.name, job.duration)
		jobs <- job
	}()

	// Render success.
	w.WriteHeader(http.StatusCreated)
	return
}

func main() {
	var (
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
		port         = flag.String("port", "8080", "The server port")
	)
	flag.Parse()

	// create job channel
	jobs := make(chan job, *maxQueueSize)

	// create workers
	for i := 1; i <= *maxWorkers; i++ {
		go func(i int) {
			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	// handler for adding jobs
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(jobs, w, r)
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

/*

$ go run worker_httpd.go -max_workers 5
added: job1 2s
worker2: started job1, working for 2.000000 seconds
added: job2 3s
worker1: started job2, working for 3.000000 seconds
added: job3 4s
worker3: started job3, working for 4.000000 seconds
added: job4 5s
worker4: started job4, working for 5.000000 seconds
added: job5 6s
worker5: started job5, working for 6.000000 seconds
added: job6 7s
added: job7 8s
added: job8 9s
added: job9 1s
added: job10 2s
added: job11 3s
added: job12 4s
added: job13 5s
added: job14 6s
added: job15 7s
worker2: completed job1!
worker2: started job6, working for 7.000000 seconds
worker1: completed job2!
worker1: started job7, working for 8.000000 seconds
worker3: completed job3!
worker3: started job8, working for 9.000000 seconds
. . .
worker3: completed job8!
worker2: completed job13!
worker5: completed job14!
worker1: completed job15!
^Csignal: interrupt

cURL the application from another terminal window after launching above:

   for i in {1..15}; do curl localhost:8080/work -d name=job$i -d delay=$(expr $i % 9 + 1)s; done

*/
