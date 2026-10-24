package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC19ToolingE2(-5); got != 16 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 16) } }
func TestHidden2(t *testing.T) { if got := TransformC19ToolingE2(0); got != 21 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 21) } }
func TestHidden3(t *testing.T) { if got := TransformC19ToolingE2(1000000); got != 1000021 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000021) } }
