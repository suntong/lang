package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of butchers is simply the length of this list.
var ph = []string{"A", "B", "C", "D", "E"}

const hunger = 3                // Number of times each butcher works
const rest = time.Second / 100 // Mean rest time
const work = time.Second / 100   // Mean work time

var fmt = log.New(os.Stdout, "", 0)

var chopping sync.WaitGroup

func choppingProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(phName, "Positioned")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		fmt.Println(phName, "Collecting")
		dominantHand.Lock() // pick up knifes
		otherHand.Lock()
		fmt.Println(phName, "Working")
		rSleep(work)
		dominantHand.Unlock() // put down knifes
		otherHand.Unlock()
		fmt.Println(phName, "Resting")
		rSleep(rest)
	}
	fmt.Println(phName, "Done")
	chopping.Done()
	fmt.Println(phName, "Left")
}

func main() {
	chopping.Add(5)
	knife0 := &sync.Mutex{}
	knifeLeft := knife0
	for i := 1; i < len(ph); i++ {
		knifeRight := &sync.Mutex{}
		go choppingProblem(ph[i], knifeLeft, knifeRight)
		knifeLeft = knifeRight
	}
	go choppingProblem(ph[0], knife0, knifeLeft)
	chopping.Wait() // wait for butchers to finish
	fmt.Println("Work place empty")
}
