package main

import "testing"

func Test1(t *testing.T) { if got := TransformC09PointersE2(1); got != 12 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 12) } }
func Test2(t *testing.T) { if got := TransformC09PointersE2(2); got != 13 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 13) } }
func Test3(t *testing.T) { if got := TransformC09PointersE2(-1); got != 10 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 10) } }
