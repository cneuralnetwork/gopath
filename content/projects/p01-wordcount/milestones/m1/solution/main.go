package main

import "fmt"

// TransformP01WordcountM1 applies the requested transformation.
func TransformP01WordcountM1(n int) int {
	return addOffset(n, 11)
}

func main() { fmt.Println(TransformP01WordcountM1(2)) }
