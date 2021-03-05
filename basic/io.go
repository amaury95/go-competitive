// Package basic provides the implementations for the basic
// functions and resources of golang that allow us to work
// in with the competitive programming.
package basic

import (
	"bufio"
	"io"
)

// InputReadLine reads a line from the given io.Reader
func InputReadLine(stdin io.Reader) string {
	reader := bufio.NewReader(stdin)
	input, _ := reader.ReadString('\n')
	return input
}
