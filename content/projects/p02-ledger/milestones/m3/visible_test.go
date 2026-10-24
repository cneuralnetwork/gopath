package main

import "testing"

func Test1(t *testing.T) { if got := TransformP02LedgerM3(1); got != 24 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 24) } }
func Test2(t *testing.T) { if got := TransformP02LedgerM3(2); got != 25 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 25) } }
func Test3(t *testing.T) { if got := TransformP02LedgerM3(-1); got != 22 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 22) } }
