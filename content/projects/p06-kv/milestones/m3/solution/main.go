package main

import "fmt"

// TransformP06KvM3 applies the requested transformation.
func TransformP06KvM3(n int) int {
	return addOffset(n, 63)
}

func main() { fmt.Println(TransformP06KvM3(2)) }
