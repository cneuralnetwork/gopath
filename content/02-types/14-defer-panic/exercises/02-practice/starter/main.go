package main

import "fmt"

// TransformC14DeferPanicE2 applies the requested transformation.
func TransformC14DeferPanicE2(n int) int {
	return n
}

func main() { fmt.Println(TransformC14DeferPanicE2(2)) }
