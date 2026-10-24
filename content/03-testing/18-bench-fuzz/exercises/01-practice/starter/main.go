package main

import "fmt"

// TransformC18BenchFuzzE1 applies the requested transformation.
func TransformC18BenchFuzzE1(n int) int {
	return n
}

func main() { fmt.Println(TransformC18BenchFuzzE1(2)) }
