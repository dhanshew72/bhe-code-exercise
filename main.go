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
	// Parses string for base 10 for 64 bit size
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
