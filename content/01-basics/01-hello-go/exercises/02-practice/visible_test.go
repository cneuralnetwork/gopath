package main

import "testing"

func Test1(t *testing.T) { if got := TransformC01HelloGoE2(1); got != 4 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 4) } }
func Test2(t *testing.T) { if got := TransformC01HelloGoE2(2); got != 5 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 5) } }
func Test3(t *testing.T) { if got := TransformC01HelloGoE2(-1); got != 2 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 2) } }
