package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC09PointersE1(-5); got != 5 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 5) } }
func TestHidden2(t *testing.T) { if got := TransformC09PointersE1(0); got != 10 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 10) } }
func TestHidden3(t *testing.T) { if got := TransformC09PointersE1(1000000); got != 1000010 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000010) } }
