package fizzbuzz

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {

	// count number for test case
	counter := 100
	for i := 1; i < counter; i++ {
		// Seed the random number generator with the current time
		rand.Seed(time.Now().UnixNano())

		// Generate a random integer between 0 and 100
		randomInt := rand.Intn(101)

		v := FizzBuzz(randomInt)
		if randomInt%3 == 0 {
			assert.Equal(t, v, "Fizz", "they should be equal")
		} else if randomInt%5 == 0 {
			assert.Equal(t, v, "Buzz", "they should be equal")
		} else if randomInt%15 == 0 {
			assert.Equal(t, v, "FizzBuzz", "they should be equal")
		} else {
			assert.Equal(t, v, v, "they should be equal")
		}
	}

}
