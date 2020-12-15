package adapters

import (
	"fmt"
	"sort"
)

func Max(list []int) int {
	m := list[0]
	for _, item := range list {
		if item > m {
			m = item
		}
	}
	return m
}

func GetDiffs(list []int) map[int]int {
	diffs := map[int]int{}
	// Copy list to avoid modifying input.
	expanded := append([]int(nil), 0, Max(list)+3)
	expanded = append(expanded, list...)
	sort.Ints(expanded)
	for idx, adapter := range expanded[1:] {
		diffs[adapter-expanded[idx]]++
	}
	return diffs
}

// Counts the number of adapter arrangements.
//
// The problem is solved using dynamic programming.  The key insight is that
// after sorting the list, if it contains x and x + 3 as consecutive elements,
// then both x and x + 3 must belong to _all_ possible arrangements.  This
// implies that every arrangement of the entire list is an arrangement of the
// list from the beginning to x, plus an arragement of the list from x + 3 to
// the end.  So, the the number of arrangements of the entire list A(list) is
// equal to the product A(list[0:index(x + 3)]) * A(list[index(x + 3):-1]).
func CountArragements(list []int) int {
	expanded := append([]int(nil), 0, Max(list)+3)
	expanded = append(expanded, list...)
	sort.Ints(expanded)
	return countArrangements(expanded, 0, len(expanded))
}

func countArrangements(list []int, i, j int) int {
	// Try to break the list into two smaller lists joined by the two elements
	// required in every arrangement.
	for idx := i + 1; idx < j-1; idx++ {
		if list[idx+1]-list[idx] == 3 {
			val := countArrangements(list, i, idx+1) * countArrangements(list, idx+1, j)
			return val
		}
	}
	// If we made it here, the list contains no such required elements.  No
	// other elements are part of every arrangement, so we try to remove them
	// all, one by one.
	subslice := append([]int(nil), list[i:j]...)
	seen := map[string]bool{fmt.Sprintf("%v", subslice): true}
	bruteForceCount(subslice, seen)
	return len(seen)
}

func bruteForceCount(list []int, seen map[string]bool) {
	for i := 1; i < len(list)-1; i++ {
		if list[i+1]-list[i-1] > 3 {
			// We can't remove i from the list.
			continue
		}
		new_list := append([]int(nil), list[0:i]...)
		new_list = append(new_list, list[i+1:len(list)]...)
		seen[fmt.Sprintf("%v", new_list)] = true
		bruteForceCount(new_list, seen)
	}
}
