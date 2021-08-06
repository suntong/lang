package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// Number of butchers is simply the length of this list.
var (
	bn = []string{"A", "B", "C", "D", "E"}
	st = make([]string, len(bn))
)

const hunger = 3                // Number of times each butcher works
const rest = time.Second / 10		// Mean rest time
const work = time.Second / 10		// Mean work time

var fmt = log.New(os.Stdout, "", 0)

var chopping sync.WaitGroup

func status(i int, status string) {
	st[i] = status
	fmt.Println(strings.Join(st, " "), "\t", bn[i], status)
}

func choppingProblem(i int, dominantHand, otherHand *sync.Mutex) {
	bName := bn[i]
	//fmt.Println("\t\t", bName, "Positioned")
	status(i, "P")
	h := fnv.New64a()
	h.Write([]byte(bName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		//fmt.Println("\t\t", bName, "Collecting")
		status(i, "C")
		dominantHand.Lock() // pick up knifes
		otherHand.Lock()
		//fmt.Println("\t\t", bName, "Working")
		status(i, "W")
		rSleep(work)
		dominantHand.Unlock() // put down knifes
		otherHand.Unlock()
		//fmt.Println("\t\t", bName, "Resting")
		status(i, "R")
		rSleep(rest)
	}
	//fmt.Println("\t\t", bName, "Done")
	chopping.Done()
	//fmt.Println("\t\t", bName, "Left")
	status(i, "-")
}

func main() {
	fmt.Println(strings.Join(bn, " "))
	fmt.Println(strings.Repeat("--", len(bn)))
	for i := 1; i < len(st); i++ {
		st[i] = " "
	}
	//fmt.Println(strings.Join(st, " "))
	chopping.Add(5)
	knife0 := &sync.Mutex{}
	knifeLeft := knife0
	for i := 1; i < len(bn); i++ {
		knifeRight := &sync.Mutex{}
		go choppingProblem(i, knifeLeft, knifeRight)
		knifeLeft = knifeRight
	}
	go choppingProblem(0, knife0, knifeLeft)
	chopping.Wait() // wait for butchers to finish
	fmt.Println("Work place empty")
}
