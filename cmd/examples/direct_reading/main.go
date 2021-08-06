package main

import (
	"github.com/sal1entx86/boolfuck"
)

func main() {
	bf := boolfuck.BoolFuck{
		Program:     "-------++-+----++--+-+-++-++---++-++---++-++++--++--++++---", // nonsense string input here
		WriteStdOut: true,
	}

	bf.Parse(bf.Program)
}
