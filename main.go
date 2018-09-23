package main

import (
	"crypto/rand"
	"flag"
	"fmt"
)

var (
	n = flag.Int("n", 4, "id length")
)

var p = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	flag.Parse()

	b := make([]byte, *n)
	rand.Read(b)
	for i := range b {
		b[i] = p[int(b[i])%len(p)]
	}

	fmt.Print(string(b))
}
