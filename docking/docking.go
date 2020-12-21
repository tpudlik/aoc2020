package docking

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Interface satified by implementations of the decoder chip in Parts 1 and 2.
type DecoderChip interface {
	// Update the value stored at the given address.
	UpdateValue(int64, int64)
	// Update the bitmask.
	UpdateMask(string) error
	// Return the sum of all values held in memory.
	MemorySum() int64
}

var memRe = regexp.MustCompile(`mem\[(?P<address>[0-9]+)\] = (?P<value>[0-9]+)`)

// Execute the program in r on the decoder chip c.
func ExecuteProgram(r io.Reader, c DecoderChip) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(cmd, "mask = ") {
			if err := c.UpdateMask(strings.TrimPrefix(cmd, "mask = ")); err != nil {
				return err
			}
		} else if strings.HasPrefix(cmd, "mem") {
			matches := memRe.FindStringSubmatch(cmd)
			if matches == nil || len(matches) < 3 {
				return fmt.Errorf("Failed to parse command %q: did not match %v", cmd, memRe)
			}
			address, err := strconv.ParseInt(matches[memRe.SubexpIndex("address")], 10, 64)
			if err != nil {
				return fmt.Errorf("Failed to parse address in %q: %v", cmd, err)
			}
			value, err := strconv.ParseInt(matches[memRe.SubexpIndex("value")], 10, 64)
			if err != nil {
				return fmt.Errorf("Failed to parse value in %q: %v", cmd, err)
			}
			c.UpdateValue(address, value)
		} else if cmd == "" {
			// Empty string: OK to have whitespace in programs.
			continue
		} else {
			// No other command types are supported
			return fmt.Errorf("Unsupported command: %q", cmd)
		}
	}
	return scanner.Err()
}

// DecoderChip implementation for Part 1.
type Computer struct {
	mask_ones   int64
	mask_zeroes int64
	memory      map[int64]int64
}

func NewComputer() *Computer {
	return &Computer{0, 0, map[int64]int64{}}
}

func (c *Computer) UpdateMask(mask string) error {
	c.mask_ones = 0
	c.mask_zeroes = 0
	for idx, r := range mask {
		switch r {
		case 'X':
			continue
		case '1':
			c.mask_ones += 1 << (35 - idx)
		case '0':
			c.mask_zeroes += 1 << (35 - idx)
		default:
			return fmt.Errorf("Unexpected character %v in mask %q", r, mask)
		}
	}
	return nil
}

func (c *Computer) UpdateValue(address, value int64) {
	c.memory[address] = (value | c.mask_ones) & (^c.mask_zeroes)
}

func (c *Computer) MemorySum() int64 {
	out := int64(0)
	for _, v := range c.memory {
		out += v
	}
	return out
}

// DecoderChip implementation for Part 2.
type MemoryAddressDecoderChip struct {
	memory        map[int64]int64
	mask_ones     int64
	mask_floating map[int64]bool
}

func NewMemoryAddressDecoderChip() *MemoryAddressDecoderChip {
	return &MemoryAddressDecoderChip{map[int64]int64{}, 0, map[int64]bool{}}
}

func (c *MemoryAddressDecoderChip) UpdateMask(mask string) error {
	c.mask_ones = 0
	c.mask_floating = map[int64]bool{}
	for idx, r := range mask {
		switch r {
		case 'X':
			c.mask_floating[int64(idx)] = true
		case '1':
			c.mask_ones += 1 << (35 - idx)
		case '0':
			continue
		default:
			return fmt.Errorf("Unexpected character %v in mask %q", r, mask)
		}
	}
	return nil
}

func (c *MemoryAddressDecoderChip) UpdateValue(address, value int64) {
	addresses := map[int64]bool{address | c.mask_ones: true}
	// For every bit in the floating mask, add a variant of the address with
	// both a 0 and a 1.
	for float := range c.mask_floating {
		existing := []int64{}
		for a := range addresses {
			existing = append(existing, a)
		}
		for _, base := range existing {
			a := base | (1 << (35 - float))
			addresses[a] = true
			a = base & (^(1 << (35 - float)))
			addresses[a] = true
		}
	}
	for a := range addresses {
		c.memory[a] = value
	}
}

func (c *MemoryAddressDecoderChip) MemorySum() int64 {
	out := int64(0)
	for _, v := range c.memory {
		out += v
	}
	return out
}
