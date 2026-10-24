package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC06SlicesE2(-5); got != 3 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 3) } }
func TestHidden2(t *testing.T) { if got := TransformC06SlicesE2(0); got != 8 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 8) } }
func TestHidden3(t *testing.T) { if got := TransformC06SlicesE2(1000000); got != 1000008 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000008) } }
