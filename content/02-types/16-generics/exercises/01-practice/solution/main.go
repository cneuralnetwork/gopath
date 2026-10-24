package main

import "fmt"

// TransformC16GenericsE1 applies the requested transformation.
func TransformC16GenericsE1(n int) int {
	return n + 17
}

func main() { fmt.Println(TransformC16GenericsE1(2)) }
