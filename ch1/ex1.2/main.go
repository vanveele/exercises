// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, l := range os.Args {
		fmt.Printf("%v %v\n", i, l)
	}
}
