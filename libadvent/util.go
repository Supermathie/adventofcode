package libadvent

import "errors"

// ChanToSlice reads in all values from a channel and returns a slice with the values
func ChanToSlice(c chan string) []string {
	vals := make([]string, 0)
	for line := range c {
		vals = append(vals, line)
	}
	return vals
}

// ChanToSliceSI reads in all values from a channel and returns a slice with the values
func ChanToSliceSI(c chan []int) [][]int {
	vals := make([][]int, 0)
	for line := range c {
		vals = append(vals, line)
	}
	return vals
}

// Sum sums all the items of data
func Sum(data []int) (total int) {
	total = 0
	for _, d := range data {
		total += d
	}
	return
}

// Min finds the minimum
func Min(data []int) (min int) {
	min = data[0]
	for _, d := range data {
		if d < min {
			min = d
		}
	}
	return
}

// Max finds the maximum
func Max(data []int) (max int) {
	max = data[0]
	for _, d := range data {
		if d > max {
			max = d
		}
	}
	return
}

// FindCombinationTotal does what you think
func FindCombinationTotal(data []int, num int, target int) []int {
	for c := range Combinations(data, num) {
		if Sum(c) == target {
			return c
		}
	}
	return nil
}

// IndexOf returns the index of needle in the haystack
func IndexOf(haystack []int, needle int) (int, error) {
	for i, v := range haystack {
		if v == needle {
			return i, nil
		}
	}
	return 0, errors.New("cannot find needle")
}

// IndexOfS returns the index of needle in the haystack
func IndexOfS(haystack []string, needle string) (int, bool) {
	for i, v := range haystack {
		if v == needle {
			return i, true
		}
	}
	return 0, false
}

// LCM returns the Least Common Multiple of A and B
func LCM(a, b uint64) uint64 {
	return a * b / GCD(a, b)
}

// GCD returns the Greatest Common Divisor of A and B
func GCD(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// AllCombinations generates all combinations of options
func AllCombinations(options []int) chan []int {
	c := make(chan []int, 4)
	go func() {
		defer close(c)
		for i := 0; i <= len(options); i++ {
			subC := Combinations(options, i)
			for combo := range subC {
				c <- combo
			}
		}
	}()
	return c
}

// Combinations generates all num-combinations of options
func Combinations(options []int, num int) chan []int {
	c := make(chan []int, 2)
	go func() {
		defer close(c)
		if num == 0 {
			c <- []int{}
		} else {
			for i, val := range options {
				if num == 1 {
					c <- []int{val}
				} else {
					for subCombination := range Combinations(options[i+1:], num-1) {
						combination := make([]int, num)
						combination[0] = val
						copy(combination[1:], subCombination)
						c <- combination
					}
				}
			}
		}
	}()
	return c
}
