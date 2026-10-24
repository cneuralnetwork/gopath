package main

import "fmt"

// TransformP02LedgerM2 applies the requested transformation.
func TransformP02LedgerM2(n int) int {
	return addOffset(n, 22)
}

func main() { fmt.Println(TransformP02LedgerM2(2)) }
