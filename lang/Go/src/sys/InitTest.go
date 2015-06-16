package main

func init() {
	println("first")
}

func init() {
	println("second")
}

func init() {
	println("next")
}

func main() {
	// even though main() is empty, this program will print
	// "first" then "second" and "next"
}
