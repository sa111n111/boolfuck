package main

import (
	"github.com/sa111n111/boolfuck"
)

func main() {
	bf := boolfuck.BoolFuck{
		WriteStdOut: true,
	}
	bf.Open("test.bf")
}
