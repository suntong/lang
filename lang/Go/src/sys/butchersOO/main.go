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

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

var (
	// butcher Names
	bNms = []string{"A", "B", "C", "D", "E"}
	// Number of butchers is simply the length of this list.
	bLen = len(bNms)
)

const hunger = 3              // Number of times each butcher works
const rest = time.Second / 10 // Mean rest time
const work = time.Second / 10 // Mean work time

var fmt = log.New(os.Stdout, "", 0)

//==========================================================================
// Status

//go:generate stringer -type=statusN -linecomment=true

// status names
type statusN int

const (
	Initialized statusN = iota // _
	Positioned                 // P
	Collecting                 // C
	Working                    // W
	Resting                    // R
	Done                       // D
	Left                       // .
)

// var statusA = []string{
// 	Initialized: " ", Positioned: "P", Collecting: "C",
// 	Resting: "R", Done: "D", Left: "-",
// }

// status values
type statusV struct {
	sync.Mutex
	st []string
}

// var st = statusV{ st: make([]string, bLen) }

func newStatusV(butchers int) statusV {
	st := statusV{st: make([]string, butchers)}
	for i := 1; i < butchers; i++ {
		// st.st[i] = statusA[Initialized]
		st.st[i] = Initialized.String()
	}
	fmt.Println("Legend:")
	fmt.Printf(" %s: %s ->\n", Positioned, "Positioned")
	fmt.Printf(" %s: %s ->\n", Collecting, "Collecting")
	fmt.Printf(" %s: %s ->\n", Resting, "Resting")
	fmt.Printf(" %s: %s ->\n", Done, "Done")
	fmt.Printf(" %s: %s\n\n", Left, "Left")
	fmt.Println(strings.Join(bNms, " "))
	fmt.Println(strings.Repeat("--", butchers))
	return st
}

func (st *statusV) statusUpdate(i int, status statusN) {
	st.Lock()
	st.st[i] = status.String()
	fmt.Println(strings.Join(st.st, " "), "\t", bNms[i], st.st[i])
	st.Unlock()
}

//==========================================================================
// chopping

type choppingActivity struct {
	sync.WaitGroup
}

func newChoppingActivity(butchers int) choppingActivity {
	chopping := choppingActivity{}
	chopping.Add(butchers)
	return chopping
}

func (chopping choppingActivity) choppingAction(
	i int, dominantHand, otherHand *sync.Mutex, st *statusV) {
	bName := bNms[i]
	st.statusUpdate(i, Positioned)
	h := fnv.New64a()
	h.Write([]byte(bName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		st.statusUpdate(i, Collecting)
		dominantHand.Lock() // pick up knifes
		otherHand.Lock()
		st.statusUpdate(i, Working)
		rSleep(work)
		dominantHand.Unlock() // put down knifes
		otherHand.Unlock()
		st.statusUpdate(i, Resting)
		rSleep(rest)
	}
	//fmt.Println("\t\t", bName, "Done")
	chopping.Done()
	st.statusUpdate(i, Left)
}

func (chopping choppingActivity) choppingSimulation(butchers int) {
	st := newStatusV(butchers)
	knife0 := &sync.Mutex{}
	knifeLeft := knife0
	for i := 1; i < bLen; i++ {
		knifeRight := &sync.Mutex{}
		go chopping.choppingAction(i, knifeLeft, knifeRight, &st)
		knifeLeft = knifeRight
	}
	go chopping.choppingAction(0, knife0, knifeLeft, &st)
	chopping.Wait() // wait for butchers to finish
	fmt.Println("Work place empty")
}

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	c := newChoppingActivity(bLen)
	c.choppingSimulation(bLen)
}
