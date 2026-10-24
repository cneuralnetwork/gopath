package main

import "fmt"

// TransformP02LedgerM3 applies the requested transformation.
func TransformP02LedgerM3(n int) int {
	return addOffset(n, 23)
}

func main() { fmt.Println(TransformP02LedgerM3(2)) }
