package main

import "fmt"

// TransformP03TodoM1 applies the requested transformation.
func TransformP03TodoM1(n int) int {
	return addOffset(n, 31)
}

func main() { fmt.Println(TransformP03TodoM1(2)) }
