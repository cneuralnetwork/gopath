package main

import "testing"

func Test1(t *testing.T) { if got := TransformC02VariablesE2(1); got != 5 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 5) } }
func Test2(t *testing.T) { if got := TransformC02VariablesE2(2); got != 6 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 6) } }
func Test3(t *testing.T) { if got := TransformC02VariablesE2(-1); got != 3 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 3) } }
