package main

import "fmt"

// TransformP05ShortenerM2 applies the requested transformation.
func TransformP05ShortenerM2(n int) int {
	return addOffset(n, 52)
}

func main() { fmt.Println(TransformP05ShortenerM2(2)) }
