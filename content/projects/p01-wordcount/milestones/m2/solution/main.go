package main

import "fmt"

// TransformP01WordcountM2 applies the requested transformation.
func TransformP01WordcountM2(n int) int {
	return addOffset(n, 12)
}

func main() { fmt.Println(TransformP01WordcountM2(2)) }
