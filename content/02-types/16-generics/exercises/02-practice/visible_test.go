package main

import "testing"

func Test1(t *testing.T) { if got := TransformC16GenericsE2(1); got != 19 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 19) } }
func Test2(t *testing.T) { if got := TransformC16GenericsE2(2); got != 20 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 20) } }
func Test3(t *testing.T) { if got := TransformC16GenericsE2(-1); got != 17 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 17) } }
