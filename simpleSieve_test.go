package fastsegmentedsieve

import (
	"reflect"
	"testing"
)

func TestSimpleSeive(t *testing.T) {
	shouldPrimes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	if !reflect.DeepEqual(shouldPrimes, SimpleSieve(40)) {
		t.Error("Expected slice of primes")
	}
}
