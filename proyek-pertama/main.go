package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	HelloWorld()
}

func HelloWorld() string {
	text := "Hello World!"
	fmt.Println(text)
	return text
}
