package sieve

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime_InvalidNegative(t *testing.T) {
	_, err := NewSieve().NthPrime(-1)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidValue.Error())
}

func TestNthPrime_0(t *testing.T) {
	actual, err := NewSieve().NthPrime(0)
	expected := int64(2)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_1(t *testing.T) {
	actual, err := NewSieve().NthPrime(1)
	expected := int64(3)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_19(t *testing.T) {
	actual, err := NewSieve().NthPrime(19)
	expected := int64(71)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_99(t *testing.T) {
	actual, err := NewSieve().NthPrime(99)
	expected := int64(541)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_500(t *testing.T) {
	actual, err := NewSieve().NthPrime(500)
	expected := int64(3_581)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_986(t *testing.T) {
	actual, err := NewSieve().NthPrime(986)
	expected := int64(7_793)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_2000(t *testing.T) {
	actual, err := NewSieve().NthPrime(2_000)
	expected := int64(17_393)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_1000000(t *testing.T) {
	actual, err := NewSieve().NthPrime(1_000_000)
	expected := int64(15_485_867)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNthPrime_10000000(t *testing.T) {
	actual, err := NewSieve().NthPrime(10_000_000)
	expected := int64(179_424_691)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

// not required, just a fun challenge
// requires segmented sieve approach, other approach hit timeout limits of 30 seconds.
func TestNthPrime_100000000(t *testing.T) {
	actual, err := NewSieve().NthPrime(100_000_000)
	expected := int64(2_038_074_751)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()
	f.Fuzz(func(t *testing.T, num int64) {
		if num < 0 {
			// Skip negative values for fuzz testing.
			// See: TestNthPrime_InvalidNegative for check
			t.Skip()
		}
		result, err := sieve.NthPrime(num)
		if err != nil {
			t.Fatalf("NthPrime(%d) returned an unexpected error: %v", num, err)
		}
		if !big.NewInt(result).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", num)
		}
	})
}
