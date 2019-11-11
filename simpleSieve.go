package fastsegmentedsieve

import "math"

// SimpleSieve returns prime numbers starting from 2 till the limit specified.
func SimpleSieve(limit int64) []int64 {
	primes := []int64{}
	if limit >= 2 {
		primes = append(primes, 2)
	}
	if limit >= 3 {
		primes = append(primes, 3)
	}
	if limit >= 5 {
		multiple := int64(limit / 6)
		if multiple*6+1 < limit {
			multiple++
		}
		size := multiple * 2
		unmark := make([]bool, size)
		var startPos int64
		for p := int64(1); 10*multiple > 6*p+1; p++ {
			num1 := 6*p - 1
			num2 := 6*p + 1
		Num1Loop:
			startPos = int64(math.Pow(10, float64(p-1))) + 7
			for i := 0; startPos+num1*2*i < size; i++ {
				unmark[startPos+num1*i] = true
			}
			startPos = p * 10
			pos := int((p - 3) / 2)
			if !unmark[pos] {
				for i := pos + int(p); i < int(size); i += int(p) {
					unmark[i] = true
				}
			}
		}
		for p := 0; p < size; p++ {
			if !unmark[p] {
				primes = append(primes, int64(p*2+3))
			}
		}
	}
	return primes
}
