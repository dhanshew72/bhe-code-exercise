// Package sieve implements the Sieve of Eratosthenes for generating prime numbers.
package sieve

import (
	"math"
)

const (
	FIRST_PRIME int64 = 2
	LIMIT_MULT  int64 = 2
)

type Sieve interface {
	NthPrime(num int64) (int64, error)
}

type solver struct{}

// NewSieve returns a new Sieve implementation.
func NewSieve() Sieve {
	return &solver{}
}

// NthPrime returns the prime at the given 0-based index, doubling the search
// limit until enough primes are found.
func (solv *solver) NthPrime(num int64) (int64, error) {
	if num < 0 {
		return 0, ErrInvalidValue
	}
	limit := FIRST_PRIME
	for {
		primes := segmentedSieve(limit)
		// If we have more primes then the input value, we have the prime value
		if int64(len(primes)) > num {
			return primes[num], nil
		}
		// This is an exponential check each iteration as sieve
		// methodology needs an upper bound
		// Without knowing the answer, we can start at 2 then
		// guess up values for the nth prime
		limit *= LIMIT_MULT
	}
}

// sieve returns all primes up to the limit
// Note: this is a brute force implementation
func sieve(limit int64) []int64 {
	basePrimes := simpleSieve(limit)
	return basePrimes
}

// segmentedSieve returns all primes up to limit using the segmented sieve method.
// It sieves a small base segment of size √limit, then processes the remaining
// range in chunks of the same size, keeping memory usage low.
func segmentedSieve(limit int64) []int64 {
	segmentSize := getSegmentSize(limit)
	basePrimes := simpleSieve(segmentSize)
	primes := append([]int64{}, basePrimes...)

	for low := segmentSize + 1; low <= limit; low += segmentSize {
		high := low + segmentSize - 1
		if high > limit {
			high = limit
		}

		segment := newSegment(low, high)
		markComposites(segment, basePrimes, low, high)

		for index, isPrime := range segment {
			if isPrime {
				primes = append(primes, low+int64(index))
			}
		}
	}

	return primes
}

// newSegment allocates a bool slice for [low, high], all true (assumed prime until marked otherwise).
func newSegment(low, high int64) []bool {
	segment := make([]bool, high-low+1)
	for i := range segment {
		segment[i] = true
	}
	return segment
}

// markComposites sieves each base prime across [low, high], marking their multiples as composite.
func markComposites(segment []bool, basePrimes []int64, low, high int64) {
	for _, prime := range basePrimes {
		// Ceiling division: first multiple of prime that falls within this segment.
		start := ((low + prime - 1) / prime) * prime
		for multiple := start; multiple <= high; multiple += prime {
			segment[multiple-low] = false
		}
	}
}

// getSegmentSize returns segment covers the boundary number exactly at sqrt(limit)
func getSegmentSize(limit int64) int64 {
	return int64(math.Sqrt(float64(limit))) + 1
}

// buildPasePrimes returns all the primes based up to a size value
func simpleSieve(limit int64) []int64 {
	primes := []int64{}
	isPrime := buildPrimeSieve(limit)
	for num := FIRST_PRIME; num <= limit; num++ {
		if isPrime[num] {
			primes = append(primes, num)
		}
	}
	return primes
}

// buildPrimeSieve returns a boolean slice of size limit+1 where index i is true if i is prime.
func buildPrimeSieve(limit int64) []bool {
	composites := make([]bool, limit+1)
	for index := range composites {
		composites[index] = true
	}
	// For each prime factor up to sqrt(limit), mark all its multiples starting at factor^2 as composite.
	// See: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	for factor := FIRST_PRIME; factor <= int64(math.Sqrt(float64(limit))); factor++ {
		if composites[factor] {
			for multiple := factor * factor; multiple <= limit; multiple += factor {
				composites[multiple] = false
			}
		}
	}
	return composites
}
