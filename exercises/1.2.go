// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.

//This exercise modifies the example text book snippet to print the command as well as the args.
package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for index, value := range os.Args {
		fmt.Printf("%d\t%s\n", index, value)
	}
}
//!-
