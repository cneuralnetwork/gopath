package main

import "testing"

func Test1(t *testing.T) { if got := TransformC28HttpE1(1); got != 30 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 30) } }
func Test2(t *testing.T) { if got := TransformC28HttpE1(2); got != 31 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 31) } }
func Test3(t *testing.T) { if got := TransformC28HttpE1(-1); got != 28 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 28) } }
