// See
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

package main

import (  
    "fmt"
    "time"
)

type Payload struct {
    data int
}

func (p *Payload) UploadToS3() error {
    fmt.Println("payload data:",p.data)
    return nil
}

var payloads = []Payload{Payload{1},Payload{2},Payload{3},Payload{4}}

func main() {
	testNOK()
	testOK2()
	testOK()
	testPointer()
}

func testNOK() {  
    fmt.Println("NOK test results:")
    for _,payload := range payloads {
        go payload.UploadToS3()
    }

    time.Sleep(1 * time.Second)
    /*
    PRINTS:
    
    payload data: 4
    payload data: 4
    payload data: 4
    payload data: 4
    */
}

func testOK2() {
	fmt.Println("OK2 test results:")
	payloads := []*Payload{{1},{2},{3},{4}}
	for _,payload := range payloads {
		go payload.UploadToS3()
   }

	time.Sleep(1 * time.Second)
}

func testOK() {  
	fmt.Println("OK test results:")
	for _,payload := range payloads {
		p := payload
		go p.UploadToS3()
    }

    time.Sleep(1 * time.Second)
}


func testPointer() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {    // 此时迭代值 v 是三个元素值的地址，每次 v 指向的值不同
		go v.print()
	}
	time.Sleep(1 * time.Second)
	// 输出 one two three
}

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}
