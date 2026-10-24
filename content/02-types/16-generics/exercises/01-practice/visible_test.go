package main

import "testing"

func Test1(t *testing.T) { if got := TransformC16GenericsE1(1); got != 18 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 18) } }
func Test2(t *testing.T) { if got := TransformC16GenericsE1(2); got != 19 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 19) } }
func Test3(t *testing.T) { if got := TransformC16GenericsE1(-1); got != 16 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 16) } }
