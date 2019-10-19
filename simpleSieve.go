package fastsegmentedsieve

// SimpleSieve returns prime numbers starting from 2 till the limit specified.
func SimpleSieve(limit int64) []int64 {
	primes := []int64{2}
	var size int
	if limit%2 == 0 {
		size = int((limit - 1) / 2)
	} else {
		size = int(limit / 2)
	}

	var unmark = make([]bool, size)
	for p := int64(3); p*p < limit; p += 2 {
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
	return primes
}
