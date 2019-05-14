package internal

import (
	"fmt"
	"io"
)

func SayHello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}
