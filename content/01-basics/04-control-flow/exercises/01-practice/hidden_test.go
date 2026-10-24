package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC04ControlFlowE1(-5); got != 0 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 0) } }
func TestHidden2(t *testing.T) { if got := TransformC04ControlFlowE1(0); got != 5 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 5) } }
func TestHidden3(t *testing.T) { if got := TransformC04ControlFlowE1(1000000); got != 1000005 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000005) } }
