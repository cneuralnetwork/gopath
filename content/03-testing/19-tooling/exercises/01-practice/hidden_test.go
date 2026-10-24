package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC19ToolingE1(-5); got != 15 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 15) } }
func TestHidden2(t *testing.T) { if got := TransformC19ToolingE1(0); got != 20 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 20) } }
func TestHidden3(t *testing.T) { if got := TransformC19ToolingE1(1000000); got != 1000020 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000020) } }
