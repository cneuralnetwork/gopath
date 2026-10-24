package main

import "fmt"

// TransformP03TodoM2 applies the requested transformation.
func TransformP03TodoM2(n int) int {
	return addOffset(n, 32)
}

func main() { fmt.Println(TransformP03TodoM2(2)) }
