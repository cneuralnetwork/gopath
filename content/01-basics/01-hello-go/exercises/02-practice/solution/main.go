package main

import "fmt"

// TransformC01HelloGoE2 applies the requested transformation.
func TransformC01HelloGoE2(n int) int {
	return n + 3
}

func main() { fmt.Println(TransformC01HelloGoE2(2)) }
