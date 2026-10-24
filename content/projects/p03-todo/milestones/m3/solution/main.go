package main

import "fmt"

// TransformP03TodoM3 applies the requested transformation.
func TransformP03TodoM3(n int) int {
	return addOffset(n, 33)
}

func main() { fmt.Println(TransformP03TodoM3(2)) }
