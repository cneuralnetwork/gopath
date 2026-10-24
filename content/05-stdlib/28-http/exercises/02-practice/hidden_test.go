package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC28HttpE2(-5); got != 25 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 25) } }
func TestHidden2(t *testing.T) { if got := TransformC28HttpE2(0); got != 30 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 30) } }
func TestHidden3(t *testing.T) { if got := TransformC28HttpE2(1000000); got != 1000030 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000030) } }
