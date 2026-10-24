package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC10StructsE2(-5); got != 7 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 7) } }
func TestHidden2(t *testing.T) { if got := TransformC10StructsE2(0); got != 12 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 12) } }
func TestHidden3(t *testing.T) { if got := TransformC10StructsE2(1000000); got != 1000012 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000012) } }
