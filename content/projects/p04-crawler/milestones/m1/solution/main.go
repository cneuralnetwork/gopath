package main

import "fmt"

// TransformP04CrawlerM1 applies the requested transformation.
func TransformP04CrawlerM1(n int) int {
	return addOffset(n, 41)
}

func main() { fmt.Println(TransformP04CrawlerM1(2)) }
