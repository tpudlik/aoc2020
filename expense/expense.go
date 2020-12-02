// Functions helpful in solving the Day 1 puzzle.
package expense

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
}

func ReadInputFromFile(fname string) ([]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadInput(file)
}

func ReadInput(r io.Reader) ([]int, error) {
	out := []int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		i, err := strconv.Atoi(txt)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode line %q as integer", txt)
		}
		out = append(out, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error scanning input file: %v", err)
	}
	return out, nil
}

func EntriesWithSum(entries []int, sum int) (int, error) {
	for i, a := range entries {
		for j, b := range entries {
			if i == j {
				continue
			}
			if a+b == sum {
				return a * b, nil
			}
		}
	}
	return 0, fmt.Errorf("No two numbers in provided slice sum up to %d", sum)
}

func ThreeEntriesWithSum(entries []int, sum int) (int, error) {
	for i, a := range entries {
		for j, b := range entries {
			if i == j {
				continue
			}
			for k, c := range entries {
				if k == i || k == j {
					continue
				}
				if a+b+c == sum {
					return a * b * c, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("No three numbers in provided slice sum up to %d", sum)
}
