package main

import (
	"crypto/rand"
	"flag"
	"fmt"
)

const defaultLength = 4

var (
	n = flag.Int("n", defaultLength, "id length")
)

var p = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	flag.Parse()

	if *n <= 0 {
		*n = defaultLength
	}

	b := make([]byte, *n)
	rand.Read(b)
	for i := range b {
		b[i] = p[int(b[i])%len(p)]
	}

	fmt.Print(string(b))
}
