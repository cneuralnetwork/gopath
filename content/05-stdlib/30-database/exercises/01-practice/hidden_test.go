package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC30DatabaseE1(-5); got != 26 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 26) } }
func TestHidden2(t *testing.T) { if got := TransformC30DatabaseE1(0); got != 31 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 31) } }
func TestHidden3(t *testing.T) { if got := TransformC30DatabaseE1(1000000); got != 1000031 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000031) } }
