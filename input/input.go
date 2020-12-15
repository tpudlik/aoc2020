package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntList(fname string) ([]int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		list = append(list, n)
	}
	return list, nil
}
