package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC10StructsE1(-5); got != 6 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 6) } }
func TestHidden2(t *testing.T) { if got := TransformC10StructsE1(0); got != 11 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 11) } }
func TestHidden3(t *testing.T) { if got := TransformC10StructsE1(1000000); got != 1000011 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000011) } }
