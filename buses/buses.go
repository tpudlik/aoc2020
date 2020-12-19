package buses

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type ParsedNotes struct {
	ts    int
	buses []int
}

func Part1(notes ParsedNotes) int {
	earliest_ts := math.MaxInt32
	earliest_bus := 0
	for _, bus := range notes.buses {
		next_arrival := (notes.ts/bus + 1) * bus
		if next_arrival < earliest_ts {
			earliest_ts = next_arrival
			earliest_bus = bus
		}
	}
	return earliest_bus * (earliest_ts - notes.ts)
}

func ParseSchedules(r io.Reader) (ParsedNotes, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	ts, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ParsedNotes{}, fmt.Errorf("Failed to parse %q as a timestamp: %v", scanner.Text(), err)
	}

	scanner.Scan()
	buses := []int{}
	for _, id := range strings.Split(scanner.Text(), ",") {
		if id == "x" {
			continue
		}
		bus, err := strconv.Atoi(id)
		if err != nil {
			return ParsedNotes{}, fmt.Errorf("Failed to parse %q as a bus id: %v", id, err)
		}
		buses = append(buses, bus)
	}
	return ParsedNotes{ts, buses}, nil
}

// Represents a set of congruences, i.e. equations of the form,
//
//   x = a mod n
type Congruences struct {
	a []int
	n []int
}

func ParseCongruences(r io.Reader) (Congruences, error) {
	scanner := bufio.NewScanner(r)

	// Skip the first line
	scanner.Scan()

	// Second line contains the congruences.
	scanner.Scan()
	as := []int{}
	ns := []int{}
	for a, n_str := range strings.Split(scanner.Text(), ",") {
		if n_str == "x" {
			continue
		}
		n, err := strconv.Atoi(n_str)
		if err != nil {
			return Congruences{}, fmt.Errorf("Failed to parse %q as an integer: %v", n_str, err)
		}
		as = append(as, n-a)
		ns = append(ns, n)
	}
	return Congruences{as, ns}, nil
}

func SolveCongruences(c Congruences) (int64, error) {
	if !ArePairwiseCoprime(c.n) {
		return 0, fmt.Errorf("The n's %v are not pairwise coprime", c.n)
	}
	a := []*big.Int{}
	for _, elem := range c.a {
		a = append(a, big.NewInt(int64(elem)))
	}
	n := []*big.Int{}
	for _, elem := range c.n {
		n = append(n, big.NewInt(int64(elem)))
	}
	return crtConstruct(a, n).Int64(), nil
}

func ArePairwiseCoprime(n []int) bool {
	for idx, n1 := range n[:len(n)-1] {
		for _, n2 := range n[idx+1:] {
			if gcd(n1, n2) > 1 {
				return false
			}
		}
	}
	return true
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

// Function below copied from
// https://golangnews.org/2020/12/computing-the-chinese-remainder-theorem/
func crtConstruct(a, n []*big.Int) *big.Int {
	// Compute N: product(n[...])
	N := new(big.Int).Set(n[0])
	for _, nk := range n[1:] {
		N.Mul(N, nk)
	}

	// x is the accumulated answer.
	x := new(big.Int)

	for k, nk := range n {
		// Nk = N/nk
		Nk := new(big.Int).Div(N, nk)

		// N'k (Nkp) is the multiplicative inverse of Nk modulo nk.
		Nkp := new(big.Int)
		if Nkp.ModInverse(Nk, nk) == nil {
			return big.NewInt(-1)
		}

		// x += ak*Nk*Nkp
		x.Add(x, Nkp.Mul(a[k], Nkp.Mul(Nkp, Nk)))
	}
	return x.Mod(x, N)
}
