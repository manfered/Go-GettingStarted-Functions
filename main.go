package main

import (
	"errors"
	"fmt"
)

func main() {
	startServer1(8000)

	result2 := startServer2(9000)
	fmt.Println(result2)

	errorResult3 := startServer3(10000)
	fmt.Println(errorResult3)

	result4, errorResult4 := startServer4(11000)
	fmt.Println(result4, errorResult4)

	// writeonly variable if we do not need to use the variable in order to prevent compilation error
	_, errorResult4_1 := startServer4(12000)
	fmt.Println(errorResult4_1)
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

// function with parameter return multiple values
func startServer4(port int) (bool, error) {
	fmt.Println("Server3 is starting...")
	fmt.Println("Server3 failed on port", port)
	return true, nil
}
