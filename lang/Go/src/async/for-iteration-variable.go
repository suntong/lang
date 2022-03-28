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
	testOK()
}

func testNOK() {  
    fmt.Println("NOK test results:")
    for _,payload := range payloads {
        go payload.UploadToS3()
    }

    time.Sleep(2 * time.Second)
    /*
    PRINTS:
    
    payload data: 4
    payload data: 4
    payload data: 4
    payload data: 4
    */
}

func testOK() {  
	fmt.Println("OK test results:")
	for _,payload := range payloads {
		p := payload
		go p.UploadToS3()
    }

    time.Sleep(2 * time.Second)
}

