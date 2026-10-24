package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC18BenchFuzzE1(-5); got != 14 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 14) } }
func TestHidden2(t *testing.T) { if got := TransformC18BenchFuzzE1(0); got != 19 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 19) } }
func TestHidden3(t *testing.T) { if got := TransformC18BenchFuzzE1(1000000); got != 1000019 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000019) } }
