package main

import (
	"errors"
	"fmt"
)

func main() {
	startServer1(8000)

	result2 := startServer2(9000)
	fmt.Println(result2)

	errorResult := startServer3(10000)
	fmt.Println(errorResult)
}

// function with parameter but no return value
func startServer1(port int) {
	fmt.Println("Server1 is starting...")
	fmt.Println("Server1 Started on port", port)
}

// function with parameter and a return value
func startServer2(port int) bool {
	fmt.Println("Server2 is starting...")
	fmt.Println("Server2 Started on port", port)
	return true
}

// function with parameter return error
func startServer3(port int) error {
	fmt.Println("Server3 is starting...")
	fmt.Println("Server3 failed on port", port)
	return errors.New("something went wrong")
}
