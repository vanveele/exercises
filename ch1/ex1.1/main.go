// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Println(strings.Join(append([]string{path.Base(os.Args[0])}, os.Args[1:]...), " "))
}
