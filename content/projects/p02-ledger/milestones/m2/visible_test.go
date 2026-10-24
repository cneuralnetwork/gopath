package main

import "testing"

func Test1(t *testing.T) { if got := TransformP02LedgerM2(1); got != 23 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 23) } }
func Test2(t *testing.T) { if got := TransformP02LedgerM2(2); got != 24 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 24) } }
func Test3(t *testing.T) { if got := TransformP02LedgerM2(-1); got != 21 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 21) } }
