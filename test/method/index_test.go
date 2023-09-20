package method

import (
	"errors"
	"testing"
)

func TestIsNegative(log *testing.T) {
	err := errors.New("Is Negative")
	if IsNegative(-1) {
		log.Log("-1 is negative")
	} else {
		log.Fatal(err)
	}

	if IsNegative(1) {
		log.Log("1 is negative")
	} else {
		log.Fatal("isNegative(1) should be false")
	}
}

func BenchmarkIsNegative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsNegative(-1)
	}
}

// go test ./test/method -v
