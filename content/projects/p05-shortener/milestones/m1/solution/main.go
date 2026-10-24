package main

import "fmt"

// TransformP05ShortenerM1 applies the requested transformation.
func TransformP05ShortenerM1(n int) int {
	return addOffset(n, 51)
}

func main() { fmt.Println(TransformP05ShortenerM1(2)) }
