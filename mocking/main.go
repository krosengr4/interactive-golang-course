package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/*
- Requirements
	1. Print 3, 2, 1, and Go!
	2. Wait one second between each line.
*/

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
