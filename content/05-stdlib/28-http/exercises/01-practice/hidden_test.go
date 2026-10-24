package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC28HttpE1(-5); got != 24 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 24) } }
func TestHidden2(t *testing.T) { if got := TransformC28HttpE1(0); got != 29 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 29) } }
func TestHidden3(t *testing.T) { if got := TransformC28HttpE1(1000000); got != 1000029 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000029) } }
