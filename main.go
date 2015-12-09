package main

import (
	"fmt"
	"os"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {

	}
	f := tempconv.Fahrenheit(23)
	fmt.Fprintf(os.Stderr, "cf: %v\n", f)
	fmt.Fprintf(os.Stderr, "Freezing temp is %v\n", tempconv.FreezingC)
	fmt.Fprintf(os.Stdout, "celsius: %v\n", tempconv.FToC(f))

}
