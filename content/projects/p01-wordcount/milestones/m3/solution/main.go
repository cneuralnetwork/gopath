package main

import "fmt"

// TransformP01WordcountM3 applies the requested transformation.
func TransformP01WordcountM3(n int) int {
	return addOffset(n, 13)
}

func main() { fmt.Println(TransformP01WordcountM3(2)) }
