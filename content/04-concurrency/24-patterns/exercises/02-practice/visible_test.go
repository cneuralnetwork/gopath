package main

import "testing"

func Test1(t *testing.T) { if got := TransformC24PatternsE2(1); got != 27 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 27) } }
func Test2(t *testing.T) { if got := TransformC24PatternsE2(2); got != 28 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 28) } }
func Test3(t *testing.T) { if got := TransformC24PatternsE2(-1); got != 25 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 25) } }
