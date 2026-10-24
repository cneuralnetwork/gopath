package main

import "testing"

func Test1(t *testing.T) { if got := TransformC06SlicesE2(1); got != 9 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 9) } }
func Test2(t *testing.T) { if got := TransformC06SlicesE2(2); got != 10 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 10) } }
func Test3(t *testing.T) { if got := TransformC06SlicesE2(-1); got != 7 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 7) } }
