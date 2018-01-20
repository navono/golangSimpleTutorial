package main

import (
	"fmt"

	_ "golangSimpleTutorial/sqlite"
	// _ "golangSimpleTutorial/formDemo"
	// _ "golangSimpleTutorial/httpServer"
	// _ "golangSimpleTutorial/concurrencyDemo"
	// _ "golangSimpleTutorial/interfaceDemo"
	// _ "golangSimpleTutorial/methodDemo"
)

func add(a int) int {
	a = a + 1
	return a
}

func main() {
	fmt.Println("================= main =================")
	fmt.Printf("Hello, Go.")
}
