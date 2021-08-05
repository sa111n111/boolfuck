package main

import "boolfuck"

func main() {
	bf := boolfuck.BoolFuck{
		Program:     "-++-+----++--+-+-++-++---++-++---++-++++",
		WriteStdOut: true,
	}

	bf.Parse()
}
