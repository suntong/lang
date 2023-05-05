package main

var printed bool

func main() {

	for i := 0; i < 3; i++ {
		println(i)
		go func() {
			println("in")
			if !printed {
				println(i)
				printed = true
			}
		}()

	}
}
