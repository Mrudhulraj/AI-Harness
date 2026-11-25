package main

import "log"

func main() {
	input, err := NewUIInput()
	if err != nil {
		log.Fatal(err)
	}
	NewUi(input)
}
