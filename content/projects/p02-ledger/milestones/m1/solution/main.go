package main

import "fmt"

// TransformP02LedgerM1 applies the requested transformation.
func TransformP02LedgerM1(n int) int {
	return addOffset(n, 21)
}

func main() { fmt.Println(TransformP02LedgerM1(2)) }
