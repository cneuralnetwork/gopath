package main

import "fmt"

// TransformC25ContextE1 applies the requested transformation.
func TransformC25ContextE1(n int) int {
	return n + 26
}

func main() { fmt.Println(TransformC25ContextE1(2)) }
