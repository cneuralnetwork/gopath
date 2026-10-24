package main

import "testing"

func Test1(t *testing.T) { if got := TransformC18BenchFuzzE2(1); got != 21 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 21) } }
func Test2(t *testing.T) { if got := TransformC18BenchFuzzE2(2); got != 22 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 22) } }
func Test3(t *testing.T) { if got := TransformC18BenchFuzzE2(-1); got != 19 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 19) } }
