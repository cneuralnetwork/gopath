package main

import "testing"

func Test1(t *testing.T) { if got := TransformC24PatternsE1(1); got != 26 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 26) } }
func Test2(t *testing.T) { if got := TransformC24PatternsE1(2); got != 27 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 27) } }
func Test3(t *testing.T) { if got := TransformC24PatternsE1(-1); got != 24 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 24) } }
