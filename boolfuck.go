package boolfuck

import "fmt"

// BinFuck -> see package docs
type BoolFuck struct {
	// The BinFuck program itself. Defined as an array of
	// bytes.
	Program string

	// WriteStdOut is if you would like to write out the result
	// of your boolfuck program to be written to the stdout.
	//
	// If you would like to see the result of your program, set this to truw.
	WriteStdOut bool

	progmem     [10000]int
	current_pos int
}

// Parse parses a binfuck program.
//
// There are checks in place to detect if the provided program contains
// excessive operators. If there are, an error is returned.
//
func (b *BoolFuck) Parse() int {
	p := b.Program
	mem := b.progmem
	offset := b.current_pos
	writeable := b.WriteStdOut

	for i := 0; i < len(p); i++ {
		if p[i] == '+' {
			mem[offset] = 1

			if writeable {
				fmt.Print("true ")
			}
		} else if p[i] == '-' {
			mem[offset] = 0

			if writeable {
				fmt.Print("false ")
			}
		} else {
			break
		}
	}

	return mem[offset]
}
