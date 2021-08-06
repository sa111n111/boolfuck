package main

import (
	"github.com/sal1entx86/boolfuck"
)

func main() {
	bf := boolfuck.BoolFuck{
		WriteStdOut: true,
	}
	bf.Open("test.bf")
}
