package main

import (
	"github.com/sa111n111/boolfuck"
)

func main() {
	bf := boolfuck.BoolFuck{
		Program:     "-------++-+----++--+-+-++-++---++-++---++-++++--++--++++---", // nonsense string input here
		WriteStdOut: true,
	}

	bf.Parse(bf.Program)
}
