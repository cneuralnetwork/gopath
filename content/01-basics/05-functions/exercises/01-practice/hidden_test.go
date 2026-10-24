package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC05FunctionsE1(-5); got != 1 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 1) } }
func TestHidden2(t *testing.T) { if got := TransformC05FunctionsE1(0); got != 6 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 6) } }
func TestHidden3(t *testing.T) { if got := TransformC05FunctionsE1(1000000); got != 1000006 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000006) } }
