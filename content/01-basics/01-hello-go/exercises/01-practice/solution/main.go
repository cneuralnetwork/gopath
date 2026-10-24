package main

import "fmt"

// TransformC01HelloGoE1 applies the requested transformation.
func TransformC01HelloGoE1(n int) int {
	return n + 2
}

func main() { fmt.Println(TransformC01HelloGoE1(2)) }
