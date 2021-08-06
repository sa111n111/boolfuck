package boolfuck

import (
	"bytes"
	"fmt"
	"os"
)

// BoolFuck is a esoteric programming language that consists of both + and - symbols
// both being representations of true or false.
type BoolFuck struct {
	// The BinFuck program itself. Defined as string
	Program string

	// WriteStdOut is if you would like to write out the result
	// of your boolfuck program to be written to the stdout.
	//
	// If you would like to see the result of your program, set this to true.
	WriteStdOut bool

	progmem     [10000]int
	current_pos int
}

// Parse parses a boolfuck program.
func (b *BoolFuck) Parse(program string) int {
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

// Open opens a boolfuck file and fills in the
// `Program` field. Parse() is then called on the read
// contents. If provided, Parse() will directly write the result
// of the read file to the stdout.
func (b *BoolFuck) Open(filename string) {
	file, err := os.Open(filename)
	buf := new(bytes.Buffer)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf.ReadFrom(file)

	contents := buf.String()
	b.Program = contents

	b.Parse(b.Program)
}
