package main

import "testing"

func Test1(t *testing.T) { if got := TransformC30DatabaseE2(1); got != 33 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 33) } }
func Test2(t *testing.T) { if got := TransformC30DatabaseE2(2); got != 34 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 34) } }
func Test3(t *testing.T) { if got := TransformC30DatabaseE2(-1); got != 31 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 31) } }
