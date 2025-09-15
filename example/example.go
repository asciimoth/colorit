package main //nolint

import (
	"os"

	"github.com/asciimoth/colorit"
)

const help = `
Usage: tool [OPTION]... [FILE]...
Tool that do something

With no FILE, or when FILE is -, read standard input.

	-A, --aaa         Aaaaa
	-B, --bbb         Bbbbb
	-c
      --help        display this help and exit

Examples:
  tool -A -B  Do something
`

const golang = `
package main

func main() {
	println("Hello World")
}
`

func main() {
	// go run example/example.go help
	// go run example/example.go golang
	switch os.Args[1] {
	case "help":
		colorit.HighlightTo(help, "help", os.Stdout) //nolint
	case "golang":
		colorit.HighlightTo(golang, "go", os.Stdout) //nolint
	}
}
