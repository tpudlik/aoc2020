package handheld

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Operation int

const (
	ACC Operation = iota
	JMP
	NOP
)

type Instruction struct {
	op  Operation
	arg int
}

type Computer struct {
	pc      int
	acc     int
	program []Instruction
}

func (c *Computer) LoadProgram(program []Instruction) {
	c.program = program
	c.pc = 0
}

// Executes the next instruction in the program.  Returns false if program end reached.
func (c *Computer) Step() bool {
	if c.pc >= len(c.program) {
		return false
	}
	instruction := c.program[c.pc]
	switch instruction.op {
	case ACC:
		c.acc += instruction.arg
		c.pc++
	case JMP:
		c.pc += instruction.arg
	case NOP:
		c.pc++
	}
	return true
}

// Detects the first repeated visit to a pc value, and returns the acc at that
// time.  Returns an error if the program terminates.
func DetectInfiniteLoop(program []Instruction) (int, error) {
	visited := map[int]bool{0: true}
	var c Computer
	c.LoadProgram(program)
	for c.Step() {
		if _, ok := visited[c.pc]; ok {
			return c.acc, nil
		}
		visited[c.pc] = true
	}
	return c.acc, fmt.Errorf("Program terminated.")
}

// Solves Day 8, Part 2.
func FixBootLoop(program []Instruction) (int, error) {
	for idx, instruction := range program {
		var new_op Operation
		switch instruction.op {
		case JMP:
			new_op = NOP
		case NOP:
			new_op = JMP
		default:
			continue
		}

		new_program := make([]Instruction, len(program))
		copy(new_program, program)
		new_program[idx] = Instruction{new_op, instruction.arg}
		acc, err := DetectInfiniteLoop(new_program)
		if err != nil {
			// Program has terminated!
			return acc, nil
		}
	}
	return 0, fmt.Errorf("No single-instruction fix worked")
}

func ParseProgram(r io.Reader) ([]Instruction, error) {
	scanner := bufio.NewScanner(r)
	var program []Instruction
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid token: %q", scanner.Text())
		}

		var op Operation
		switch tokens[0] {
		case "acc":
			op = ACC
		case "jmp":
			op = JMP
		case "nop":
			op = NOP
		default:
			return nil, fmt.Errorf("Invalid operation code: %q", tokens[0])
		}

		arg, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}

		program = append(program, Instruction{op, arg})
	}
	return program, nil
}
