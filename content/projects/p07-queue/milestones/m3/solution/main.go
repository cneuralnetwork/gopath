package main

import "fmt"

// TransformP07QueueM3 applies the requested transformation.
func TransformP07QueueM3(n int) int {
	return addOffset(n, 73)
}

func main() { fmt.Println(TransformP07QueueM3(2)) }
