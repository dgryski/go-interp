// Package interp is an interpolation search
package interp

// SearchInts searches the array for the key, returning the index of the first occurrence of the element.
func SearchInts(array []int, key int) int {
	return Search(len(array), array[0], array[len(array)-1], key, func(i int) int { return array[i] })
}

// modified from http://data.linkedin.com/blog/2010/06/beating-binary-search

// Search finds the lowest value of i such that keyAt(i) = key or keyAt(i+1) > key.
func Search(n, min, max, key int, keyAt func(i int) int) int {

	low, high := 0, n-1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		element := keyAt(guess)
		if element == key {
			// scan backwards for start of value range
			for guess > 0 && keyAt(guess-1) == key {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if element > key {
			high = guess - 1
			max = keyAt(high)
		} else {
			low = guess + 1
			min = keyAt(low)
		}
	}
}
