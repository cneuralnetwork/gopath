package main

import "testing"

func Test1(t *testing.T) { if got := TransformC01HelloGoE1(1); got != 3 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 3) } }
func Test2(t *testing.T) { if got := TransformC01HelloGoE1(2); got != 4 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 4) } }
func Test3(t *testing.T) { if got := TransformC01HelloGoE1(-1); got != 1 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 1) } }
