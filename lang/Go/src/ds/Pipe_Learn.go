////////////////////////////////////////////////////////////////////////////
// Porgram: Pipe_Learn.go
// Purpose: io.Pipe demo from Go语言学习园地
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.1.md
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	Pipe()
}

func Pipe() {
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(1e7)
}

func PipeWrite(pipeWriter *io.PipeWriter) {
	var (
		i   = 0
		err error
		n   int
	)
	data := []byte("Go语言学习园地")
	// 循环往管道中写数据，写第三次时，我们调用 CloseWithError 方法关闭管道的写入端，之后再一次调用 Write 方法，发现返回了error，于是退出了循环。
	for _, err = pipeWriter.Write(data); err == nil; n, err = pipeWriter.Write(data) {
		i++
		if i == 3 {
			// 对于管道的close方法（非CloseWithError时），err会被置为EOF
			pipeWriter.CloseWithError(errors.New("输出3次后结束"))
		}
	}
	fmt.Println("close 后输出的字节数：", n, " error：", err)
}

func PipeRead(pipeReader *io.PipeReader) {
	var (
		err error
		n   int
	)
	data := make([]byte, 1024)
	for n, err = pipeReader.Read(data); err == nil; n, err = pipeReader.Read(data) {
		fmt.Printf("%s\n", data[:n])
	}
	fmt.Println("writer 端 closewitherror后：", err)
}

/*

Output :

Go语言学习园地
Go语言学习园地
Go语言学习园地
close 后输出的字节数： 0  error： io: read/write on closed pipe
writer 端 closewitherror后： 输出3次后结束

*/
