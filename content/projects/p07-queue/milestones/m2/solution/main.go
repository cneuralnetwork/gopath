package main

import "fmt"

// TransformP07QueueM2 applies the requested transformation.
func TransformP07QueueM2(n int) int {
	return addOffset(n, 72)
}

func main() { fmt.Println(TransformP07QueueM2(2)) }
