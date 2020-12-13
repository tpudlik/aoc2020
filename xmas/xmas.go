package xmas

import "fmt"

// Returns true if n can be expressed as a sum of two different integers in
// the list.
func ExpressibleAsSum(list []int, n int) bool {
	for i, x := range list[:len(list)-1] {
		for _, y := range list[i+1:] {
			if x+y == n {
				return true
			}
		}
	}
	return false
}

func FirstInvalidNumber(list []int, preamble_length int) (int, error) {
	if len(list) < preamble_length {
		return 0, fmt.Errorf("List of numbers %v shorter than preamble length %d", list, preamble_length)
	}

	preamble := list[0:preamble_length]
	for i := preamble_length; i < len(list); i++ {
		if !ExpressibleAsSum(preamble, list[i]) {
			return list[i], nil
		}
		preamble = list[i-preamble_length : i+1]
	}
	return 0, fmt.Errorf("No invalid number in list")
}

func ContiguousSetSummingTo(list []int, n int) ([]int, error) {
	for first := 0; first < len(list)-1; first++ {
		for length := 2; first+length <= len(list); length++ {
			sum := 0
			for i := first; i < first+length; i++ {
				sum += list[i]
			}
			if sum == n {
				return list[first : first+length], nil
			}
			if sum > n {
				// Consecutive sum is already too large, and will only grow larger.
				// Advance first number.
				break
			}
		}
	}
	return nil, fmt.Errorf("No contiguous set summing to %d found in list", n)
}
