package main

import (
	"bytes"
	"fmt"
	"github.com/tbeckham/stringutil"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

// Add 2 numbers
func add(x, y int) int {
	return x + y
}

// Returns 2 strings!
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Efficient string concat using buffer
	var buffer bytes.Buffer
	buffer.WriteString(stringutil.Reverse((" ,olleH")))
	buffer.WriteString(stringutil.Reverse(("!oG")))
	fmt.Println(buffer.String())

	// The environment in which these programs are executed is deterministic, so
	// rand.Intn will always return the same number.
	//
	// To resolve this we create seed pool for each run in order to initialize the
	// default Source to a deterministic state each run
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))

	// Some function calls.
	fmt.Println("I added 42 + 27 to get", add(42, 27))
	fmt.Println(swap("Tony", "Beckham"))

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
