package main

import "fmt"

// TransformC18BenchFuzzE1 applies the requested transformation.
func TransformC18BenchFuzzE1(n int) int {
	return n + 19
}

func main() { fmt.Println(TransformC18BenchFuzzE1(2)) }
