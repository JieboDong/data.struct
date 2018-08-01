package main

import (
	"fmt"
)

func main() {

	data := make(chan string)
	go func() {

		data <- "协程"

	}()
	go func() {

		data <- "协程2"
	}()
	go func() {

		data <- "协程3"
	}()
	go func() {

		data <- "协程4"
	}()

	str := "hello world"
	fmt.Println(str)
	fmt.Println(len(<-data))

}
