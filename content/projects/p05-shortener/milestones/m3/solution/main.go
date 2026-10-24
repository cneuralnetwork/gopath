package main

import "fmt"

// TransformP05ShortenerM3 applies the requested transformation.
func TransformP05ShortenerM3(n int) int {
	return addOffset(n, 53)
}

func main() { fmt.Println(TransformP05ShortenerM3(2)) }
