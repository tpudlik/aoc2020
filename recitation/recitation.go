package recitation

// Returns the number spoken after the given number of steps.
func Recite(starting []int, steps int) int {
	// Maps each number to the step at which it was last spoken (updated with a one-step delay).
	last_spoken := map[int]int{}
	// Most recently spoken number.
	var n int
	// The second-most-recently spoken number.
	pren := -1
	for step := 0; step < steps; step++ {
		if step < len(starting) {
			n = starting[step]
		} else {
			// When was the last time the previous number was spoken?
			previous, ok := last_spoken[n]
			if !ok {
				previous = step - 1
			}
			n = (step - 1) - previous
		}
		last_spoken[pren] = step - 1
		pren = n
	}
	return n
}

// Slow reference implementation of Recite.
func ReciteRef(starting []int, steps int) int {
	spoken := []int{}
	var n int
	for step := 0; step < steps; step++ {
		if step < len(starting) {
			n = starting[step]
		} else {
			last := spoken[step-1]

			// When was the last time this number was spoken?
			previous := step - 1
			for idx := len(spoken) - 2; idx >= 0; idx-- {
				if spoken[idx] == last {
					previous = idx
					break
				}
			}
			n = (step - 1) - previous
		}
		spoken = append(spoken, n)
	}
	return n
}
