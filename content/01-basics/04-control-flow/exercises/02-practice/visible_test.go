package main

import "testing"

func Test1(t *testing.T) { if got := TransformC04ControlFlowE2(1); got != 7 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 7) } }
func Test2(t *testing.T) { if got := TransformC04ControlFlowE2(2); got != 8 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 8) } }
func Test3(t *testing.T) { if got := TransformC04ControlFlowE2(-1); got != 5 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 5) } }
