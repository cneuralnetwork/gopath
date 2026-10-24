package main

import "fmt"

// TransformP06KvM1 applies the requested transformation.
func TransformP06KvM1(n int) int {
	return addOffset(n, 61)
}

func main() { fmt.Println(TransformP06KvM1(2)) }
