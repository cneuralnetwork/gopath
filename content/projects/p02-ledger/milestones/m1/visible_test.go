package main

import "testing"

func Test1(t *testing.T) { if got := TransformP02LedgerM1(1); got != 22 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 22) } }
func Test2(t *testing.T) { if got := TransformP02LedgerM1(2); got != 23 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 23) } }
func Test3(t *testing.T) { if got := TransformP02LedgerM1(-1); got != 20 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 20) } }
