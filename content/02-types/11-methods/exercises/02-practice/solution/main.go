package main

import "fmt"

// TransformC11MethodsE2 applies the requested transformation.
func TransformC11MethodsE2(n int) int {
	return n + 13
}

func main() { fmt.Println(TransformC11MethodsE2(2)) }
