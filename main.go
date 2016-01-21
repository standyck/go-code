package main

import (
	"fmt"
	"net/http"
	"os"

	"gopl.io/ch2/popcount"
	"gopl.io/ch2/tempconv"
)

var HTTPClient *http.Client

func arrays() {
	/*	ab := [...]int{1, 2, 3} */
	a := [3]int{1, 2}
	fmt.Println(a[0])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}

func othercrap() {
	f := tempconv.Fahrenheit(23)
	fmt.Fprintf(os.Stderr, "cf: %v\n", f)
	fmt.Fprintf(os.Stderr, "Freezing temp is %v\n", tempconv.FreezingC)
	fmt.Fprintf(os.Stdout, "celsius: %v\n", tempconv.FToC(f))
	HTTPClient = http.DefaultClient
	req, err := http.NewRequest("GET", "http://www.standyck.com", nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "err: %v\n", err)
	} else {
		r, _ := HTTPClient.Do(req)
		fmt.Fprintf(os.Stdout, "req: %v\n", req)
		fmt.Fprintf(os.Stdout, "response: %v\n", r)
	}

	z := popcount.PopCount(19235)
	fmt.Printf("popcount: %v\n", z)
}

func main() {
	arrays()
}
