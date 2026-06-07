// Command sieve finds the Nth prime number (0-based index).
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"ssse-exercise-sieve/pkg/sieve"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: sieve <n>")
	}
	num, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	result, err := sieve.NewSieve().NthPrime(num)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(result)
}
