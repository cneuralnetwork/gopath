package main

import "testing"

func Test1(t *testing.T) { if got := TransformC06SlicesE1(1); got != 8 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 8) } }
func Test2(t *testing.T) { if got := TransformC06SlicesE1(2); got != 9 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 9) } }
func Test3(t *testing.T) { if got := TransformC06SlicesE1(-1); got != 6 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 6) } }
