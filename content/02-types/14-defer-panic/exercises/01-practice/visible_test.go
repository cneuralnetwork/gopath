package main

import "testing"

func Test1(t *testing.T) { if got := TransformC14DeferPanicE1(1); got != 16 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 16) } }
func Test2(t *testing.T) { if got := TransformC14DeferPanicE1(2); got != 17 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 17) } }
func Test3(t *testing.T) { if got := TransformC14DeferPanicE1(-1); got != 14 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 14) } }
