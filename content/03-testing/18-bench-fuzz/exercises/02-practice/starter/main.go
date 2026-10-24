package main

import "fmt"

// TransformC18BenchFuzzE2 applies the requested transformation.
func TransformC18BenchFuzzE2(n int) int {
	return n
}

func main() { fmt.Println(TransformC18BenchFuzzE2(2)) }
